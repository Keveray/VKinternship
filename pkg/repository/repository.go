package repository

type Authorization interface {
}

type VkInternshipList interface {
}

type VkInternshipItem interface {
}

type Repository struct {
	Authorization
	VkInternshipList
	VkInternshipItem
}

func NewRepository() *Repository {
	return &Repository{}
}
