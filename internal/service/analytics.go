
package service

import "meeting-rooms/internal/models"

func (ms *MeetingService) GetAnalytics() ([]*models.Analytics, error) {
 analytics, err := ms.Store.GetAnalytics()
 if err != nil {
  return nil, err
 }

 return analytics, nil
}
