package dbops

import (
	"database/sql"
	"log"
	"time"
	"video_server/api/defs"
	"video_server/api/ntils"
	_"video_server/api/ntils"
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
	vid, err := ntils.NewUUID()
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
	stmOut, err := dbConn.Prepare("select * from video_info where id=?")

	var aid int
	var dct string
	var name string

	err = stmOut.QueryRow(vid).Scan(&aid, &name, &dct)
	//此处返回如果没数据 也是有返回值的
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	defer stmOut.Close()

	res := &defs.VideoInfo{Id:vid, AuthorId:aid, Name:name, DisplayCtime:dct}

	return res, nil
}
//删除视频
func DeleteVideoInfo(vid string) error {
	stmtDel, err := dbConn.Prepare("DELETE from video_info WHERE id=?")
	if err != nil {
		return  err
	}
	_, err = stmtDel.Exec(vid)
	if err != nil {
		return err
	}
	defer stmtDel.Close()
	return nil
}
//添加评论
func AddNewComments(vid string, aid int, content string) error {
	id, err := ntils.NewUUID()
	if err != nil {
		return err
	}
	stmtIns, err := dbConn.Prepare(`INSERT INTO comments (id, video_id, author_id, content) values (?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	
	_, err = stmtIns.Exec(id, vid, aid, content)
	if err != nil {
		return  err
	}
	defer stmtIns.Close()
	return nil
}
//评论列表
func ListComments(vid string, from, to int) ([]*defs.Comment, error) {
	stmtOut, err := dbConn.Prepare(`SELECT comments.id, users.Login_name, comments.content FROM comments
		INNER JOIN users ON comments.author_id = users.id
		WHERE comments.video_id = ? AND comments.time > FROM_UNIXTIME(?) AND comments.time <= FROM_UNIXTIME(?)`)

	var res []*defs.Comment

	rows, err := stmtOut.Query(vid, from, to)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		var id, name, content string
		if err := rows.Scan(&id, &name, &content); err != nil {
			return res, err
		}
		c := &defs.Comment{Id:id, VideoId:vid, Author:name, Content:content}
		res = append(res, c)
	}
	defer stmtOut.Close()
	return res, nil
}
