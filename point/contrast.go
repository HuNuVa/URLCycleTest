package point

import (
	_ "URLCycleTest/logout"
	"fmt"
)

//定义方法,找出调用本方法的Slipoint对象中,比教出后者比前者中多出来的连接
func (a Slipoint) SliContrast(b Slipoint) string {

	var s string

	for _, pb := range b {
		for _, pa := range a {
			if pa.Url == pb.Url {

				if len(pa.DiffLink(pb.Link)) != 0 {
					fmt.Println(pa.Name, ":", pa.Url)
					s = s + "\n" + pa.Name+":"+pa.Url
				}

				for _, v := range pa.DiffLink(pb.Link) {
					if v != "" {
						fmt.Println("\t", v)
						s = s + "\n"+ " " + v
					}

				}

			}
		}
	}
	return s
}
