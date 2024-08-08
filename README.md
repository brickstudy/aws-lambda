
# AWS Lambda

This project is AWS lambda code.

<br>

## Introduction

- data catalog : ensures that all data uploaded to your S3 buckets complies with predefined data catalog rules.
  - 참고 : [AWS lambda를 활용한 S3 적재된 데이터 카탈로그 확인 도구 개발기](https://brickstudy.tistory.com/11)

<br>

![image](https://github.com/user-attachments/assets/49d5d9a9-67a0-452f-944a-64fc94947c0c)



<br>


## Installation

### 1/ Clone the repository && requirements.txt

```bash
git clone https://github.com/brickstudy/aws-lambda.git

cd aws-lambda
```


### 2/ config setting

- 경로 : ./data-catalog/.env

```sh
# MySQL
DB_USERNAME=
DB_PASSWORD=
DB_HOSTNAME=
DB_PORT=
DB_NAME=

# Discord
DISCORD_WEBHOOK_URL=
```

<br>

## Contributing
We welcome contributions to BricksAssistant!

To contribute, please follow these steps:

1. Fork the repository.
2. Create a new branch with your feature or bugfix.
3. Make your changes and commit them with clear messages.
4. Push your changes to your fork.
5. Open a pull request to the main repository.
6. Please ensure your code adheres to the project's coding standards and includes relevant tests.


