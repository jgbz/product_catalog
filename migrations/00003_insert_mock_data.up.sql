DO $$
DECLARE
    campaign1 INT;
    campaign2 INT;
    campaign3 INT;
    campaign4 INT;

BEGIN

    INSERT INTO content (ean, score, description, url, client_id) VALUES
    ('12317890123', 85, 'Description for product 1', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('23178901234', 90, 'Description for product 2', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('31789012345', 75, 'Description for product 3', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('17890123110', 95, 'Description for product 4', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('56789012317', 60, 'Description for product 5', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('67890123178', 80, 'Description for product 6', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('78901231789', 70, 'Description for product 7', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('89012317890', 88, 'Description for product 8', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('90123178901', 92, 'Description for product 9', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('01231789012', 78, 'Description for product 10', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('12317890134', 85, 'Description for product 11', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('23178901245', 91, 'Description for product 12', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('31789012356', 77, 'Description for product 13', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('17890123467', 94, 'Description for product 14', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('56789012345', 63, 'Description for product 15', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('67890123189', 82, 'Description for product 16', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('78901231790', 69, 'Description for product 17', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('89012317901', 89, 'Description for product 18', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('90123178912', 93, 'Description for product 19', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123),
    ('01231789123', 79, 'Description for product 20', 'https://media.istockphoto.com/id/1409666528/pt/foto/pet-activity-cute-puppy-dog-border-collie-lying-down-on-grass-chewing-on-stick-pet-dog-with.webp?s=2048x2048&w=is&k=20&c=-L362A3k9a7bnLf4pBLxmQ2Znic2L5FUA2X53JEBjQQ=', 123);


    INSERT INTO campaign (name) VALUES ('campaign 1') 
    RETURNING campaign_id INTO campaign1;

    INSERT INTO campaign (name) VALUES ('campaign 2') 
    RETURNING campaign_id INTO campaign2;

    INSERT INTO campaign (name) VALUES ('campaign 3') 
    RETURNING campaign_id INTO campaign3;

    INSERT INTO campaign (name) VALUES ('campaign 4') 
    RETURNING campaign_id INTO campaign4;


    INSERT INTO content (ean, score, description, url, campaign_id) VALUES
    ('9876543210123', 88, 'Campaign description for product 1', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign1),
    ('8765432101234', 92, 'Campaign description for product 2', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign1),
    ('7654321092345', 74, 'Campaign description for product 3', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign1),
    ('6543210983122', 96, 'Campaign description for product 4', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign1),
    ('5432109871732', 65, 'Campaign description for product 5', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign1),
    ('4321098765678', 83, 'Campaign description for product 6', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign2),
    ('3210987656789', 72, 'Campaign description for product 7', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign2),
    ('2109876547890', 89, 'Campaign description for product 8', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign2),
    ('1098765438901', 91, 'Campaign description for product 9', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign2),
    ('0987654329012', 76, 'Campaign description for product 10', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign2),
    ('9876543210134', 87, 'Campaign description for product 11', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign3),
    ('8765432101245', 93, 'Campaign description for product 12', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign3),
    ('7654321092356', 78, 'Campaign description for product 13', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign3),
    ('6543210983467', 94, 'Campaign description for product 14', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign3),
    ('5432109874578', 64, 'Campaign description for product 15', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign3),
    ('4321098765689', 82, 'Campaign description for product 16', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign4),
    ('3210987656790', 68, 'Campaign description for product 17', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign4),
    ('2109876547901', 90, 'Campaign description for product 18', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign4),
    ('1098765438912', 95, 'Campaign description for product 19', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign4),
    ('0987654329123', 80, 'Campaign description for product 20', 'https://media.istockphoto.com/id/1361394182/pt/foto/funny-british-shorthair-cat-portrait-looking-shocked-or-surprised.webp?s=2048x2048&w=is&k=20&c=RK9_GbQSqdrDG0Wt-4vivxCTZmnx6vRqclTbbzrXyYU=', campaign4);
    
END $$;
