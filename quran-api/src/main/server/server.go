package server

import (
	"github.com/jinzhu/gorm"
	"github.com/muslim-teachings/common/quran"
	"golang.org/x/net/context"
)

// QuranServer service implementation of QuranService
type QuranServer struct {
	DB *gorm.DB
}

// GetQuran is implementation of GetQuran grpc proto of quran.proto
func (s *QuranServer) GetQuran(ctx context.Context, emptyPayload *quran.EmptyPayload) (*quran.Quran, error) {
	return nil, nil
}

// GetAyah is implementation of GetAyah grpc proto of quran.proto
func (s *QuranServer) GetAyah(ctx context.Context, getAyahRequest *quran.GetAyahRequest) (*quran.Ayah, error) {
	return nil, nil
}

// GetSurah is implementation of GetSurah grpc proto of quran.proto
func (s *QuranServer) GetSurah(ctx context.Context, getSurahRequest *quran.GetSurahRequest) (*quran.Surah, error) {
	return nil, nil
}

// GetChapter is implementation of GetChapter grpc proto of quran.proto
func (s *QuranServer) GetChapter(ctx context.Context, getChapterRequest *quran.GetChapterRequest) (*quran.Chapter, error) {
	return nil, nil
}

// AddTafseer is implementation of AddTafseer grpc proto of quran.proto
func (s *QuranServer) AddTafseer(ctx context.Context, addTafseerRequest *quran.AddTafseerRequest) (*quran.DefaultResponse, error) {
	return nil, nil
}

// AddRecitatonToSurahStream is implementation of AddRecitatonToSurahStream grpc proto of quran.proto
func (s *QuranServer) AddRecitatonToSurahStream(stream quran.QuranService_AddRecitatonToSurahStreamServer) error {
	return nil
}

// AddRecitatonToSurah is implementation of AddRecitatonToSurah grpc proto of quran.proto
func (s *QuranServer) AddRecitatonToSurah(ctx context.Context, addRecitatonToSurahRequest *quran.AddRecitatonToSurahRequest) (*quran.DefaultResponse, error) {
	return nil, nil
}
