# Applied Concurrency in Go
This is the repository for the LinkedIn Learning course Applied Concurrency in Go. The full course is available from [LinkedIn Learning][lil-course-url].

![Applied Concurrency in Go][lil-thumbnail-url] 

Concurrency can be a confusing and intimidating topic to engineers, but it is an essential tool when writing production code as it allows you to write faster and more efficient solutions. In this course, Adelina Simion demystifies the intimidating topic of concurrency and showcases how to use the powerful tools of goroutines and channels. Go is designed with concurrency in mind so every developer should feel confident to use these powerful tools in their daily work. Join Adelina in this course to gain a thorough understanding of Go concurrency and learn how to apply it to solve some common engineering problems.

## Instructions
This repository has branches for each of the videos in the course. You can use the branch pop up menu in github to switch to a specific branch and take a look at the course at that stage, or you can add `/tree/BRANCH_NAME` to the URL to go to the branch you want to access.

## Branches
The branches are structured to correspond to the videos in the course. The naming convention is `CHAPTER#_MOVIE#`. As an example, the branch named `02_03` corresponds to the second chapter and the third video in that chapter. 
Some branches will have a beginning and an end state. These are marked with the letters `b` for "beginning" and `e` for "end". The `b` branch contains the code as it is at the beginning of the movie. The `e` branch contains the code as it is at the end of the movie. The `main` branch holds the final state of the code when in the course.

When switching from one exercise files branch to the next after making changes to the files, you may get a message like this:

    error: Your local changes to the following files would be overwritten by checkout:        [files]
    Please commit your changes or stash them before you switch branches.
    Aborting

To resolve this issue:
	
    Add changes to git using this command: git add .
	Commit changes using this command: git commit -m "some message"

## Installing
1. To use these exercise files, you must have the following installed:
	- [Go development tools](https://go.dev/doc/install)
    - [Visual Studio Code](https://code.visualstudio.com/) or any other IDE of your choice
    - [Postman](https://www.postman.com/)
    - [Git](https://git-scm.com/)
2. Clone this repository into your local machine using the terminal (Mac), CMD (Windows), or a GUI tool like SourceTree.
3. Run `go get .` from the root directory where you cloned this repository to install dependencies used in your current branch. Note, different branches have different dependencies, so make sure to run this command when you change branch.

### Instructor

Adelina Simion 
                            


                            

Check out my other courses on [LinkedIn Learning](https://www.linkedin.com/learning/instructors/adelina-simion).

[lil-course-url]: https://www.linkedin.com/learning/applied-concurrency-in-go
[lil-thumbnail-url]: https://cdn.lynda.com/course/3164282/3164282-1643050323318-16x9.jpg





