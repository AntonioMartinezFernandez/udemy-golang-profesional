package main

import "fmt"

func main() {
	/*
	 One to one
	*/
	tesla := User{
		name:   "Nikola",
		email:  "n.tesla@gmail.com",
		active: true,
	}

	einstein := User{
		name:   "Albert",
		email:  "a.einstein@gmail.com",
		active: true,
	}

	einsteinStudent := Student{
		user: einstein,
		code: "83645937",
	}

	fmt.Println(tesla)
	fmt.Println(einsteinStudent.user.email)

	/*
	 One to many
	*/
	golangCourse := Course{
		title: "Golang course",
	}

	introVideo := Video{
		title:    "Introduction",
		duration: 3,
		course:   golangCourse,
	}

	syntaxVideo := Video{
		title:    "Syntax",
		duration: 5,
		course:   golangCourse,
	}

	golangCourse.videos = []Video{introVideo, syntaxVideo}

	fmt.Println(introVideo.course.title)

	for _, video := range golangCourse.videos {
		fmt.Println(video.title)
	}
}

// One to one relation
type User struct {
	name   string
	email  string
	active bool
}

type Student struct {
	user User
	code string
}

// One to many relation
type Course struct {
	title  string
	videos []Video
}

type Video struct {
	title    string
	duration int
	course   Course
}
