insert into Users values("chicks0@fda.gov","Craig","Hicks","chicks0","McHCQujL");
insert into Users values("cporter1@dot.gov","Clarence","Porter","cporter1","Bq9VzVpkBGh");
insert into Users values("nrobertson2@imgur.com","Nancy","Robertson","nrobertson2","7XIrwd3YEtus");
insert into Users values("cmarshall3@businessweek.com","Carlos","Marshall","cmarshall3","oCV5FYVblHie");
insert into Users values("afowler4@studiopress.com","Ann","Fowler","afowler4","eVLa8v");
insert into Users values("tshaw5@1688.com","Terry","Shaw","tshaw5","ghNM3iG1g8u");
insert into Users values("adixon6@barnesandnoble.com","Andrew","Dixon","adixon6","erLhEBk");
insert into Users values("cmartin7@army.mil","Cheryl","Martin","cmartin7","cxQmcKW");
insert into Users values("wsullivan8@economist.com","Willie","Sullivan","wsullivan8","h7RdGQw");
insert into Users values("gwashington9@amazon.co.jp","Gregory","Washington","gwashington9","6F4qLqFHz7");
insert into Users values("hramireza@istockphoto.com","Howard","Ramirez","hramireza","9qxyEN3");
insert into Users values("abrownb@google.com.au","Amanda","Brown","abrownb","i5cirpO04JJ");
insert into Users values("dmoorec@e-recht24.de","Debra","Moore","dmoorec","6aCpoaKY4Rp");
insert into Users values("cjohnstond@google.es","Cheryl","Johnston","cjohnstond","AznV4jw9saV");
insert into Users values("jhayese@census.gov","Juan","Hayes","jhayese","xWaHdi");

  INSERT INTO Groups
      (`groupid`, `groupname`)
  VALUES
      (1, 'Cogibox'),
      (2, 'Podcat'),
      (3, 'Bubbletube'),
      (4, 'Skivee')
  ;

  INSERT INTO Rewards
      (`rewardid`, `groupid`, `rewardname`, `pointcost`, `rewarddescription`)
  VALUES
      (1, 2, 'DQ Blizzard', 150, 'Any Small Blizzard at Dairy Queen!'),
      (2, 2, 'T-Shirt', 500, 'Google T-Shirt'),
      (3, 2, '20 Dollars', 1000, 'MONEY MONEY MONEY!'),
      (4, 1, 'Pizza Party', 2000, 'Pizza for Everyone!'),
      (5, 2, 'Coffee Mug', 300, 'Plain Coffee Mug with Company Logo'),
      (6, 3, 'Pizza Party', 500, 'PIZZA! YAY!'),
      (7, 3, 'Coffee Mug', 200, 'Car Coffee Mug'),
      (8, 4, '20 Dollars', 1500, '20 dollars to spend on whatever you''d like!')
  ;

  INSERT INTO Tasks
      (`taskid`, `groupid`, `taskname`, `taskdescription`, `datedue`, `approvalstatus`, `completionstatus`, `pointvalue`)
  VALUES
      (1, 1, 'Get Information', 'Example Description', '2016-01-18 18:08:00', 'Approved', 'Incomplete', 400),
      (2, 4, 'Chores', 'Example Description', '2016-10-04 14:59:00', 'Pending', 'Complete', 400),
      (3, 2, 'Fax Paper', 'Example Description', '2015-10-28 09:47:00', 'Pending', 'Complete', 250),
      (4, 1, 'Fill Out Worksheet', 'Example Description', '2016-05-09 20:09:00', 'Approved', 'Incomplete', 450),
      (5, 2, 'Dishes', 'Example Description', '2015-11-16 04:22:00', 'Approved', 'Complete', 300),
      (6, 2, 'Take Out Trash', 'Example Description', '2016-09-03 16:01:00', 'Approved', 'Complete', 100),
      (7, 4, 'Fill Out Worksheet', 'Example Description', '2016-09-19 13:14:00', 'Pending', 'Complete', 250),
      (8, 3, 'Clean Car', 'Example Description', '2016-01-13 09:27:00', 'Pending', 'Complete', 200),
      (9, 4, 'Fill Out Worksheet', 'Example Description', '2016-08-08 08:52:00', 'Pending', 'Incomplete', 50),
      (10, 3, 'Dishes', 'Example Description', '2016-02-12 11:44:00', 'Approved', 'Complete', 200),
      (11, 3, 'Chores', 'Example Description', '2016-05-08 20:42:00', 'Pending', 'Complete', 150),
      (12, 4, 'Fill Out Worksheet', 'Example Description', '2016-01-31 00:58:00', 'Approved', 'Incomplete', 100),
      (13, 2, 'Fill Out Worksheet', 'Example Description', '2015-10-27 00:19:00', 'Pending', 'Complete', 400),
      (14, 4, 'Chores', 'Example Description', '2016-09-25 23:10:00', 'Pending', 'Incomplete', 200),
      (15, 4, 'Write Interview Questions', 'Example Description', '2016-05-03 08:01:00', 'Approved', 'Incomplete', 300)
  ;

  INSERT INTO TaskLeaders
      (`taskleaderid`, `emailaddress`, `groupid`)
  VALUES
      (1, 'chicks0@fda.gov', 4),
      (2, 'jhayese@census.gov', 1),
      (3, 'jhayese@census.gov', 2),
      (4, 'gwashington9@amazon.co.jp', 3),
      (5, 'jhayese@census.gov', 4)
  ;

  INSERT INTO Points
      (`pointid`, `totalpoints`, `emailaddress`, `groupid`)
  VALUES
      (1, 2500, 'chicks0@fda.gov', 4),
      (2, 225, 'cporter1@dot.gov', 1),
      (3, 1406, 'cmarshall3@businessweek.com', 3),
      (4, 3486, 'afowler4@studiopress.com', 3),
      (5, 21, 'tshaw5@1688.com', 2),
      (6, 1732, 'adixon6@barnesandnoble.com', 4),
      (7, 2538, 'cmartin7@army.mil', 4),
      (8, 161, 'wsullivan8@economist.com', 1),
      (9, 1279, 'gwashington9@amazon.co.jp', 3),
      (10, 2494, 'hramireza@istockphoto.com', 2),
      (11, 254, 'abrownb@google.com.au', 2),
      (12, 1519, 'dmoorec@e-recht24.de', 2),
      (13, 3153, 'cjohnstond@google.es', 4),
      (14, 1899, 'jhayese@census.gov', 4),
      (15, 2788, 'chicks0@fda.gov', 1),
      (16, 432, 'jhayese@census.gov', 1),
      (17, 6546, 'jhayese@census.gov', 2),
      (18, 543, 'wsullivan8@economist.com', 2),
      (19, 654, 'hramireza@istockphoto.com', 4),
      (20, 1646, 'cmarshall3@businessweek.com', 1)
  ;
