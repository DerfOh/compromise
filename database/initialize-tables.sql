CREATE TABLE Users
(
EmailAddress varchar(255),
FirstName varchar(255),
LastName varchar(255),
Nickname varchar(255),
Password varchar(255),
PRIMARY KEY(EmailAddress)
);


CREATE TABLE Groups
(
GroupId int AUTO_INCREMENT,
GroupName varchar(255),
PRIMARY KEY (GroupId)
);


CREATE TABLE Rewards
(
RewardId int AUTO_INCREMENT,
GroupId int,
RewardName varchar(255),
PointCost int,
RewardDescription varchar(255),
PRIMARY KEY (RewardId),
CONSTRAINT fk_GroupReward FOREIGN KEY (GroupId)
REFERENCES Groups (GroupId)
);


CREATE TABLE RewardRequests
(
RequestId int AUTO_INCREMENT,
RewardId int,
RewardedUser varchar(255),
PRIMARY KEY (RequestId),
CONSTRAINT fk_RewardRequest FOREIGN KEY (RewardId)
REFERENCES Rewards (RewardId),
CONSTRAINT fk_RewardUser FOREIGN KEY (RewardedUser)
REFERENCES Users (EmailAddress)
);


CREATE TABLE Tasks
(
TaskId int AUTO_INCREMENT,
GroupId int,
TaskName varchar(255),
TaskDescription varchar(255),
CompletionStatus varchar(255),
CompletedBy varchar(255),
PointValue int,
PRIMARY KEY (TaskId),
CONSTRAINT fk_UserReward FOREIGN KEY (CompletedBy)
REFERENCES Users (EmailAddress),
CONSTRAINT fk_GroupTask FOREIGN KEY (GroupId)
REFERENCES Groups(GroupId)
);


CREATE TABLE TaskLeaders
(
TaskLeaderId int AUTO_INCREMENT,
EmailAddress varchar(255),
GroupId int,
PRIMARY KEY (TaskLeaderId),
CONSTRAINT fk_GroupLeader FOREIGN KEY (GroupId)
REFERENCES Groups (GroupId),
CONSTRAINT fk_GroupUser FOREIGN KEY (EmailAddress)
REFERENCES Users (EmailAddress)
);


CREATE TABLE Points
(
PointId int AUTO_INCREMENT,
TotalPoints int DEFAULT 0,
EmailAddress varchar(255),
GroupId int,
PRIMARY KEY (PointId),
CONSTRAINT fk_GroupPoint FOREIGN KEY (GroupId)
REFERENCES Groups (GroupId),
CONSTRAINT fk_UserPoint FOREIGN KEY (EmailAddress)
REFERENCES Users (EmailAddress)
);