# Impensa

An expense tracker app for University of Queensland INFS3202 Web Information System project.

## Services

| Service     | Route  | Code Loc    | Stack                                              |
| :---------- | :----- | :---------- | :------------------------------------------------- |
| Backend API | `/api` | `/`         | [BeeGo](https://beego.me/), a Golang web framework |
| Frontend    | `/`    | `/frontend` | [React](https://reactjs.org/) with hooks           |

| Others     | Stack                                                                                                                                                                                |
| :--------- | :----------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| Database   | [Atlas](https://www.mongodb.com/cloud/atlas), a cloud MongoDB service                                                                                                                |
| Deployment | Automatic CI/CD pipeline with [GitLab CI](https://docs.gitlab.com/ee/ci/) and [Docker Compose](https://docs.docker.com/compose/) to [AWS EC2](https://aws.amazon.com/ec2/) instance. |
