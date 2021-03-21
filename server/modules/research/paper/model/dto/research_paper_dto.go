package dto



import (
"gin-vue-admin/modules/research/paper/model"
)

type ResearchPaperResponse struct {
	ResearchPaper model.ResearchPaper `json:"researchPaper"`
}


type ResearchPaperDto struct {
	Title    string `json:"title"`
}