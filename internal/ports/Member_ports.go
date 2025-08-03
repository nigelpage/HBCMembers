package ports
import (
    "/internal/domain"
)
type MemberRepository interface {
    Save(member *domain.Member) error
    FindByID(id int) (*domain.Member, error)
    FindAll() ([]*domain.Member, error)
}
