SELECT Points.GroupId, Groups.GroupName, Points.TotalPoints FROM Points JOIN Groups ON Points.GroupId = Groups.GroupId WHERE EmailAddress = 'jhayese@census.gov'; #Find Groups By Email Address
SELECT Users.EmailAddress, Users.FirstName, Users.LastName, Users.Nickname, Points.TotalPoints FROM Users JOIN Points ON Users.EmailAddress = Points.EmailAddress JOIN Groups ON Groups.GroupId = Points.GroupId WHERE Groups.GroupId = '3'; #Get Group Members Based On Group Id
SELECT * FROM Tasks WHERE GroupId = '3'; #Get Tasks By Group ID
SELECT * FROM Rewards WHERE GroupId = '2'; #Get Rewards By Group ID
SELECT * FROM TaskLeaders WHERE EmailAddress = 'chicks0@fda.gov' AND GroupId = '4'; #Find If Task Leader