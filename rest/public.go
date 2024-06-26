package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/gobkc/fit/conf"
	"github.com/gobkc/fit/driver"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	Error int    `json:"error"`
	More  string `json:"more,omitempty"`
	Msg   string `json:"msg"`
}

type VersionResponse struct {
	ApiVersion string `json:"apiVersion"`
}

// Version
//
//	@Tags		public apis
//	@Summary	Get current api version
//	@Produce	json
//	@Success	200	{object}	VersionResponse	"success"
//	@Router		/p/version [get]
func (s *Server) Version(c *gin.Context) {
	resp := VersionResponse{
		ApiVersion: s.c.Version,
	}
	c.JSON(http.StatusOK, resp)
}

// HealthCheck
//
//	@Tags		public apis
//	@Summary	k8s health check
//	@Produce	json
//	@Success	200	{string}	string	"Success"
//	@Failure	401	{object}	string	"Unauthorized"
//	@Router		/p/health [get]
func (s *Server) HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "Success")
}

type NewNoteRequest struct {
	Cate    string `json:"cate"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type NewNoteResponse struct {
	Response
	Parameters NewNoteRequest `json:"parameters"`
}

// NewNote
//
//	@Tags		public apis
//	@Summary	Add a new note
//	@Produce	json
//	@Param		data	body		NewNoteRequest	true	"request parameters, must be fill in"
//	@Success	200		{object}	NewNoteResponse	"success"
//	@Failure	401		{object}	string			"Unauthorized"
//	@Router		/p/new-note [post]
func (s *Server) NewNote(c *gin.Context) {
	request := NewNoteRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, Response{Error: 1, More: err.Error(), Msg: "ParameterError"})
		return
	}
	if err := s.d.NewNote(request.Cate, request.Title, request.Content); err != nil {
		s.JSON(c, NewNoteResponse{
			Parameters: request,
			Response: Response{
				Error: 1,
				Msg:   "Failed to create note",
				More:  err.Error(),
			},
		})
		return
	}
	s.JSON(c, NewNoteResponse{
		Parameters: request,
		Response: Response{
			Msg: "OK",
		}},
	)
}

type ListCateResponse struct {
	Response
	Data []string `json:"data"`
}

// ListCate
//
//	@Tags		public apis
//	@Summary	List notebook categories
//	@Produce	json
//	@Success	200	{object}	ListCateResponse	"success"
//	@Failure	401	{object}	string				"Unauthorized"
//	@Router		/p/list-cate [get]
func (s *Server) ListCate(c *gin.Context) {
	categories, err := s.d.ListCate()
	if err != nil {
		s.JSON(c, ListCateResponse{
			Response: Response{
				Error: 1,
				Msg:   "Failed to list categories",
				More:  err.Error(),
			},
		})
		return
	}
	s.JSON(c, ListCateResponse{
		Data: categories,
		Response: Response{
			Msg: "OK",
		}},
	)
}

type NewCateRequest struct {
	Cate string `json:"cate"`
}

type NewCateResponse struct {
	Response
	Parameters NewCateRequest `json:"parameters"`
}

// NewCate
//
//	@Tags		public apis
//	@Summary	Add a new cate
//	@Produce	json
//	@Param		data	body		NewCateRequest	true	"request parameters, must be fill in"
//	@Success	200		{object}	NewCateResponse	"success"
//	@Failure	401		{object}	string			"Unauthorized"
//	@Router		/p/new-cate [post]
func (s *Server) NewCate(c *gin.Context) {
	request := NewCateRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, Response{Error: 1, More: err.Error(), Msg: "ParameterError"})
		return
	}
	if err := s.d.NewCate(request.Cate); err != nil {
		s.JSON(c, NewCateResponse{
			Parameters: request,
			Response: Response{
				Error: 1,
				Msg:   "Failed to create category",
				More:  err.Error(),
			},
		})
		return
	}
	s.JSON(c, NewCateResponse{
		Parameters: request,
		Response: Response{
			Msg: "OK",
		}},
	)
}

type DeleteCateRequest struct {
	Cate string `json:"cate"`
}

type DeleteCateResponse struct {
	Response
	Parameters DeleteCateRequest `json:"parameters"`
}

// DeleteCate
//
//	@Tags		public apis
//	@Summary	Delete a cate
//	@Produce	json
//	@Param		data	body		DeleteCateRequest	true	"request parameters, must be fill in"
//	@Success	200		{object}	DeleteCateResponse	"success"
//	@Failure	401		{object}	string			"Unauthorized"
//	@Router		/p/cate [delete]
func (s *Server) DeleteCate(c *gin.Context) {
	request := DeleteCateRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, Response{Error: 1, More: err.Error(), Msg: "ParameterError"})
		return
	}
	if err := s.d.DeleteCate(request.Cate); err != nil {
		s.JSON(c, DeleteCateResponse{
			Parameters: request,
			Response: Response{
				Error: 1,
				Msg:   "Failed to delete a category",
				More:  err.Error(),
			},
		})
		return
	}
	s.JSON(c, DeleteCateResponse{
		Parameters: request,
		Response: Response{
			Msg: "OK",
		}},
	)
}

type DeleteNoteRequest struct {
	Cate  string `json:"cate"`
	Title string `json:"title"`
}

type DeleteNoteResponse struct {
	Response
	Parameters DeleteNoteRequest `json:"parameters"`
}

// DeleteNote
//
//	@Tags		public apis
//	@Summary	Delete a note
//	@Produce	json
//	@Param		data	body		DeleteNoteRequest	true	"request parameters, must be fill in"
//	@Success	200		{object}	DeleteNoteResponse	"success"
//	@Failure	401		{object}	string			"Unauthorized"
//	@Router		/p/note [delete]
func (s *Server) DeleteNote(c *gin.Context) {
	request := DeleteNoteRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, Response{Error: 1, More: err.Error(), Msg: "ParameterError"})
		return
	}
	if err := s.d.DeleteNote(request.Cate, request.Title); err != nil {
		s.JSON(c, DeleteNoteResponse{
			Parameters: request,
			Response: Response{
				Error: 1,
				Msg:   "Failed to delete a note",
				More:  err.Error(),
			},
		})
		return
	}
	s.JSON(c, DeleteNoteResponse{
		Parameters: request,
		Response: Response{
			Msg: "OK",
		}},
	)
}

type ListNoteResponse struct {
	Response
	Data []driver.NoteInstance `json:"data"`
}

// ListAllNote
//
//	@Tags		public apis
//	@Summary	List all notebooks
//	@Produce	json
//	@Param		keyword	query		string				false	"Keyword"
//	@Success	200		{object}	ListNoteResponse	"success"
//	@Failure	401		{object}	string				"Unauthorized"
//	@Router		/p/list-note [get]
func (s *Server) ListAllNote(c *gin.Context) {
	// this is a fake feature, list all notes using the ListNote function
}

// ListNote
//
//	@Tags		public apis
//	@Summary	List notebooks
//	@Produce	json
//	@Param		cate	path		string				false	"Category"
//	@Param		keyword	query		string				false	"Keyword"
//	@Success	200		{object}	ListNoteResponse	"success"
//	@Failure	401		{object}	string				"Unauthorized"
//	@Router		/p/{cate}/list-note [get]
func (s *Server) ListNote(c *gin.Context) {
	cate := c.Param("cate")
	keyword := c.DefaultQuery(`keyword`, ``)
	notes, err := s.d.ListNotes(cate)
	if err != nil {
		s.JSON(c, ListNoteResponse{
			Response: Response{
				Error: 1,
				Msg:   "Failed to list notes",
				More:  err.Error(),
			},
		})
		return
	}
	if keyword != `` {
		var filteredNotes []driver.NoteInstance
		split := strings.Split(keyword, ` `)
		for _, note := range notes {
			for _, key := range split {
				if strings.Contains(note.Title, key) || strings.Contains(note.Content, key) {
					filteredNotes = append(filteredNotes, note)
					break
				}
			}
		}
		notes = filteredNotes
	}
	s.JSON(c, ListNoteResponse{
		Data: notes,
		Response: Response{
			Msg: "OK",
		}},
	)
}

// Push
//
//	@Tags		public apis
//	@Summary	push all notes to email
//	@Produce	json
//	@Success	200	{object}	Response	"success"
//	@Failure	401	{object}	string		"Unauthorized"
//	@Router		/p/push [post]
func (s *Server) Push(c *gin.Context) {
	//bytes := s.d.AddFiles(s.c.JwtSalt)
	bytes := s.d.AddFiles(``)
	if err := s.d.SendEmail(s.c.Email.User, `fit`, `fit-update`, bytes); err != nil {
		s.JSON(c, Response{
			Error: 1,
			More:  err.Error(),
			Msg:   `failed to send email`,
		})
		return
	}
	s.JSON(c, Response{Msg: `ok`})
}

// Pull
//
//	@Tags		public apis
//	@Summary	pull fit attachment from email
//	@Produce	json
//	@Success	200	{object}	Response	"success"
//	@Failure	401	{object}	string		"Unauthorized"
//	@Router		/p/pull [post]
func (s *Server) Pull(c *gin.Context) {
	data, err := s.d.GetAttachmentFromEmail()
	if err != nil {
		s.JSON(c, Response{
			Error: 1,
			More:  err.Error(),
			Msg:   `failed to send email`,
		})
		return
	}
	//data, _ := os.ReadFile(`/home/xiong/fit.fit`)
	//files := s.d.DeCompress(s.c.JwtSalt, data)
	files := s.d.DeCompress(``, data)
	cachePath := conf.GetCachePath()
	if err := os.RemoveAll(cachePath); err != nil {
		s.JSON(c, Response{
			Error: 1,
			More:  err.Error(),
			Msg:   `failed to remove cache path`,
		})
		return
	}
	for _, file := range files {
		cate := strings.ReplaceAll(file.Cate, cachePath, ``)
		fileName := strings.ReplaceAll(file.Filename, `.md`, ``)
		if err := s.d.NewNote(cate, fileName, file.Content, file.UpdatedTime); err != nil {
			s.JSON(c, Response{
				Error: 1,
				Msg:   "Failed to create note",
				More:  err.Error(),
			})
			return
		}
	}
	s.JSON(c, Response{Msg: `ok`})
}

type ListConfResponse struct {
	Response
	Data     []conf.Conf `json:"data"`
	MainConf string      `json:"main_conf"`
}

// ListConf
//
//	@Tags		public apis
//	@Summary	List notebook configurations
//	@Produce	json
//	@Success	200	{object}	ListConfResponse	"success"
//	@Failure	401	{object}	string				"Unauthorized"
//	@Router		/p/list-conf [get]
func (s *Server) ListConf(c *gin.Context) {
	list, mainConf, err := s.d.ListConfigurations()
	if err != nil {
		s.JSON(c, ListConfResponse{
			Response: Response{
				Error: 1,
				Msg:   "Failed to list configurations",
				More:  err.Error(),
			},
		})
		return
	}
	s.JSON(c, ListConfResponse{
		Response: Response{
			Msg: "OK",
		},
		Data:     list,
		MainConf: mainConf,
	},
	)
}

type CreateConfRequest struct {
	Conf conf.Conf `json:"conf"`
}

type CreateConfResponse struct {
	Response
	Parameters CreateConfRequest `json:"parameters"`
}

// CreateConf
//
//	@Tags		public apis
//	@Summary	create a new configuration
//	@Produce	json
//	@Param		data	body		CreateConfRequest	true	"request parameters, must be fill in"
//	@Success	200		{object}	CreateConfResponse	"success"
//	@Failure	401		{object}	string				"Unauthorized"
//	@Router		/p/create-conf [post]
func (s *Server) CreateConf(c *gin.Context) {
	request := CreateConfRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, Response{Error: 1, More: err.Error(), Msg: "ParameterError"})
		return
	}
	if err := s.d.CreateConfiguration(request.Conf); err != nil {
		s.JSON(c, CreateConfResponse{
			Parameters: request,
			Response: Response{
				Error: 1,
				Msg:   "Failed to create configuration",
				More:  err.Error(),
			},
		})
		return
	}
	s.JSON(c, CreateConfResponse{
		Parameters: request,
		Response: Response{
			Msg: "OK",
		}},
	)
}

type EnableConfRequest struct {
	Conf conf.Conf `json:"conf"`
}

type EnableConfResponse struct {
	Response
	Parameters EnableConfRequest `json:"parameters"`
}

// EnableConf
//
//	@Tags		public apis
//	@Summary	upsert & enable a configuration
//	@Produce	json
//	@Param		data	body		EnableConfRequest	true	"request parameters, must be fill in"
//	@Success	200		{object}	EnableConfResponse	"success"
//	@Failure	401		{object}	string				"Unauthorized"
//	@Router		/p/enable-conf [post]
func (s *Server) EnableConf(c *gin.Context) {
	request := EnableConfRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, Response{Error: 1, More: err.Error(), Msg: "ParameterError"})
		return
	}
	if err := s.d.EnableConfiguration(request.Conf); err != nil {
		s.JSON(c, EnableConfResponse{
			Parameters: request,
			Response: Response{
				Error: 1,
				Msg:   "Failed to enable configuration",
				More:  err.Error(),
			},
		})
		return
	}
	s.JSON(c, EnableConfResponse{
		Parameters: request,
		Response: Response{
			Msg: "OK",
		}},
	)
}

type DeleteConfRequest struct {
	Name string `json:"name"`
}

type DeleteConfResponse struct {
	Response
	Parameters DeleteConfRequest `json:"parameters"`
}

// DeleteConf
//
//	@Tags		public apis
//	@Summary	delete a configuration
//	@Produce	json
//	@Param		data	body		DeleteConfRequest	true	"request parameters, must be fill in"
//	@Success	200		{object}	DeleteConfResponse	"success"
//	@Failure	401		{object}	string				"Unauthorized"
//	@Router		/p/conf [delete]
func (s *Server) DeleteConf(c *gin.Context) {
	request := DeleteConfRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusOK, Response{Error: 1, More: err.Error(), Msg: "ParameterError"})
		return
	}
	if err := s.d.DeleteConfiguration(request.Name); err != nil {
		s.JSON(c, DeleteConfResponse{
			Parameters: request,
			Response: Response{
				Error: 1,
				Msg:   "Failed to delete configuration",
				More:  err.Error(),
			},
		})
		return
	}
	s.JSON(c, DeleteConfResponse{
		Parameters: request,
		Response: Response{
			Msg: "OK",
		}},
	)
}
