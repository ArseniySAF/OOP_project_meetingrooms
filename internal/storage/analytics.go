package storage

import "meeting-rooms/internal/models"

func (ps *Postgres) GetAnalytics() ([]*models.Analytics, error) {
 analytics := make([]*models.Analytics, 0)

 query := `
  WITH booking_stats AS (
   SELECT
    room_id,
    COUNT(*) AS total_bookings,
    SUM(EXTRACT(EPOCH FROM (end_time - start_time)) / 3600) AS total_hours
   FROM bookings
   WHERE status IN ('confirmed', 'completed')
   GROUP BY room_id
  ),
  busiest_days AS (
   SELECT DISTINCT ON (room_id)
    room_id,
    DATE(start_time) AS busiest_day
   FROM bookings
   WHERE status IN ('confirmed', 'completed')
   GROUP BY room_id, DATE(start_time)
   ORDER BY room_id, COUNT(*) DESC
  )
  SELECT
   r.id AS room_id,
   r.name AS room_name,
   COALESCE(s.total_bookings, 0) AS total_bookings,
   COALESCE(s.total_hours, 0) AS total_hours,
   bd.busiest_day
  FROM rooms r
  LEFT JOIN booking_stats s ON s.room_id = r.id
  LEFT JOIN busiest_days bd ON bd.room_id = r.id
  WHERE r.is_active = true
  ORDER BY total_bookings DESC;
 `

 rows, err := ps.db.Query(query)

 if err != nil {
  return nil, err
 }
 defer rows.Close()

 for rows.Next() {
  analytic := &models.Analytics{}

  if err := rows.Scan(
   &analytic.RoomId,
   &analytic.RoomName,
   &analytic.TotalBookings,
   &analytic.TotalHours,
   &analytic.BusiestDay,
  ); err != nil {
   return nil, err
  }

  analytics = append(analytics, analytic)
 }

 if err := rows.Err(); err != nil {
  return nil, err
 }

 return analytics, nil
}
