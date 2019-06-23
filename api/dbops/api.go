package dbops

import (
	"database/sql"
	"log"
	"time"
	"video_server/api/defs"
	"video_server_1_5/api/utils"
)
//创建用户
func AddUserCredential(loginName string, pwd string) error {
	stmtIns, err := dbConn.Prepare("insert into users (login_name, pwd) values (?, ?)")
	if err != nil {
		return err
	}

	_, err = stmtIns.Exec(loginName, pwd)
	if err != nil {
		return err
	}
	defer stmtIns.Close()
	return nil

}
//获取密码
func GetUserCredential(loginName string) (string, error) {
	stmtOut, err := dbConn.Prepare("select pwd from users where login_name = ?")
	if err != nil {
		log.Printf("%s", err)
		return "", err
	}
	var pwd string
	err = stmtOut.QueryRow(loginName).Scan(&pwd)
	if err != nil && err != sql.ErrNoRows {
		return "", err
	}
	defer stmtOut.Close()
	return pwd, nil

}

func DeleteUser(loginName string, pwd string) error {
	stmtDel, err := dbConn.Prepare("delete from users where login_name = ? and pwd = ?")

	if err != nil {
		log.Printf("DelecteUser error %s", err)
		return err
	}
	_, err = stmtDel.Exec(loginName, pwd)
	if err != nil {
		log.Printf("DelecteUser error2 %s", err)
		return err
	}
	stmtDel.Exec()
	defer stmtDel.Close()
	return nil
}

//添加视频
func AddNewVideo(aid int, name string) (*defs.VideoInfo, error) {
	//创建 uuid
	vid, err := utils.NewUUID()
	if err != nil {
		return nil, err
	}
	t := time.Now()
	ctime := t.Format("Jan 02 2006, 15:04:05")

	stmtIns, err := dbConn.Prepare(`INSERT INTO video_info (id, author_id, name, display_ctime) values (?, ?, ?, ?)`)
	if err != nil {
		return nil, err
	}
	_, err = stmtIns.Exec(vid, aid, name, ctime)
	if err != nil {
		return nil, err
	}
	res := &defs.VideoInfo{Id:vid, AuthorId:aid, Name:name, DisplayCtime:ctime}
	defer stmtIns.Close()
	return res,nil
}
//获取视频
func GetVideoInfo(vid string) (*defs.VideoInfo, error){
	stmOut, err := dbConn.Prepare("select")

}