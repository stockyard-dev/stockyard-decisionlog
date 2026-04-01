package server
import("encoding/json";"net/http";"strconv";"github.com/stockyard-dev/stockyard-decisionlog/internal/store")
func(s *Server)handleList(w http.ResponseWriter,r *http.Request){q:=r.URL.Query().Get("q");list,_:=s.db.List(q);if list==nil{list=[]store.Decision{}};writeJSON(w,200,list)}
func(s *Server)handleCreate(w http.ResponseWriter,r *http.Request){var d store.Decision;json.NewDecoder(r.Body).Decode(&d);if d.Title==""{writeError(w,400,"title required");return};s.db.Create(&d);writeJSON(w,201,d)}
func(s *Server)handleUpdate(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);var d store.Decision;json.NewDecoder(r.Body).Decode(&d);d.ID=id;s.db.Update(&d);writeJSON(w,200,d)}
func(s *Server)handleDelete(w http.ResponseWriter,r *http.Request){id,_:=strconv.ParseInt(r.PathValue("id"),10,64);s.db.Delete(id);writeJSON(w,200,map[string]string{"status":"deleted"})}
func(s *Server)handleOverview(w http.ResponseWriter,r *http.Request){m,_:=s.db.Stats();writeJSON(w,200,m)}
