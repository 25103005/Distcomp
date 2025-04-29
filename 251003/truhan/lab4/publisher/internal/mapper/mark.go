package mapper

import (
	mark "github.com/Khmelov/Distcomp/251003/truhan/labX/publisher/internal/model"
	"github.com/Khmelov/Distcomp/251003/truhan/labX/publisher/internal/storage/model"
)

func MapMarkToModel(i mark.Mark) model.Mark {
	return model.Mark{
		ID:   int64(i.ID),
		Name: i.Name,
	}
}

func MapModelToMark(i model.Mark) mark.Mark {
	return mark.Mark{
		ID:   i.ID,
		Name: i.Name,
	}
}
