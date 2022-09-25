DROP TABLE IF EXISTS employee;
CREATE TABLE employee (
    id int NOT NULL AUTO_INCREMENT PRIMARY KEY,
    firstName VARCHAR(128) NOT NULL,
    lastName VARCHAR(128) NOT NULL,
    sex VARCHAR(10) NOT NULL,
    email VARCHAR(128) NOT NULL,
    username VARCHAR(128) NOT NULL,
    designation VARCHAR(128) NOT NULL,
    contactNumber BIGINT NOT NULL,
    bloodGroup VARCHAR(3) NULL,
    teamName VARCHAR(200) NULL,
    dateOfJoining date NOT NULL,
    home VARCHAR(128) NULL
);
INSERT INTO employee (
        firstName,
        lastName,
        sex,
        email,
        username,
        designation,
        contactNumber,
        bloodGroup,
        teamName,
        dateOfJoining,
        home
    )
VALUES (
        "Prakash",
        "Anand",
        "Male",
        "prakash.anand@maersk.com",
        "pan142",
        "Backend Developer",
        9113185872,
        "AB+",
        "New Product and Enhancement",
        "2020-10-08",
        "Marathalli Bangalore"
    ),
    (
        "Vikash",
        "Kumar",
        "Male",
        "vikash.kumar@refactor.com",
        "vkr173",
        "Frontend Developer",
        9060472145,
        "B+",
        "UI Development",
        "2020-08-26",
        "Indiranagar Bangalore"
    ),
    (
        "Aayush",
        "Somani",
        "Male",
        "aayush.somani@flipkart.com",
        "asf928",
        "Backend Developer",
        9034401789,
        "O+",
        "Backend Development",
        "2022-06-08",
        "Embassy Tech Village Bangalore"
    ),
    (
        "Meghna",
        "Singh",
        "Female",
        "meghna.singh@maersk.com",
        "mgs820",
        "SAP Developer",
        9673401546,
        "A+",
        "Customer Resolution",
        "2020-11-08",
        "Naggwara Bangalore"
    );