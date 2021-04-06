package dto



import (
"gin-vue-admin/modules/research/paper/model"
	"time"
)

type ResearchPaperResponse struct {
	ResearchPaper model.ResearchPaper `json:"researchPaper"`
}


type ResearchPaperDto struct {
	Title    string `json:"title"`
	Kind    int `json:"kind"`
	PublicationType int `json:"publicationType"`
	PublicationName string `json:"publicationName"`
	Period string `json:"period"`
	VolumeNo string `json:"volumeNo"`
	PublishingDate time.Time `json:"publishingDate"`
	WordLength       int `json:"wordLength"`
	KnowledgeClass string `json:"knowledgeClass"`
	FirstKnowledge string `json:"firstKnowledge"`
	Source string `json:"source"`
	PublishingRange string `json:"publishingRange"`
	Issn string `json:"issn"`
	Cn string `json:"cn"`
}