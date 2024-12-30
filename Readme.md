## Steps to run the application locally

- Open the cloned repo in your terminal
- Spin up RabbitMQ in a docker container locally
  ```
  docker-compose up
  ```
- Run the application
  ```
  go run main.go
  ```
- Access RabbitMQ Management Console: http://localhost:15672/

- Output:

  <img width="504" alt="image" src="https://github.com/user-attachments/assets/106c5ff7-1a72-43d6-a8db-6ce501a5074f" />

[References]:
* https://www.svix.com/resources/guides/rabbitmq-docker-setup-guide/
* https://www.rabbitmq.com/tutorials/tutorial-one-go
