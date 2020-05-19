package main

import "fmt"

import "database/sql"

import _ "mysql-master"

type daftar_buku struct{
  ID string
  Judul string
  Pengarang string
  Tahun int
}

func koneksi()(*sql.DB, error){
  db, err := sql.Open("mysql","root:@tcp(127.0.0.1:3306)/db_daftarbuku")

  if err != nil{
    return nil,err
  }
  return db,nil
}

func sql_tampil(){
  db, err := koneksi()
  if err != nil{
    fmt.Println(err.Error())
    return
  }
  defer db.Close()

  rows,err := db.Query("select * from tbl_buku")

  if err != nil{
    fmt.Println(err.Error())
    return
  }
  defer rows.Close()

  var result[]daftar_buku

  for rows.Next(){
    var each = daftar_buku{}
    var err = rows.Scan(&each.ID,&each.Judul,&each.Pengarang,&each.Tahun)

    if err != nil{
      fmt.Println(err.Error())
      return
    }
    result = append(result,each)
  }

  if err = rows.Err();err != nil{
    fmt.Println(err.Error())
    return
  }
  for _, each := range result{
    fmt.Println(each.Judul,each.Pengarang,each.Tahun)
  }
}

func main(){
 sql_tampil()
}
