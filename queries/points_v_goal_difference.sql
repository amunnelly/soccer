select team as Team,
sum(Points) as Points,
sum(GoalDifference) as GD,
sum(GoalsFor) as GoalsFor,
sum(GoalsAgainst) as GoalsAgainst
from points
where season = '%s'
and Competition = 'Premier League'
and Country = 'England'
group by team
order by points desc;
/*
 Middlesboro
 Middlesbrough
 
 Sheffield Weds
 */
update points set Team = 'Middlesbrough' where Team = 'Middlesboro';
update points set Opposition = 'Middlesbrough' where Opposition = 'Middlesboro';

update points set Team = 'Notts Forest' where Team = 'Nott''m Forest';
update points set Opposition = 'Notts Forest' where Opposition = 'Nott''m Forest';
update points set Team = 'Sheffield Wednesday' where Team = 'Sheffield Weds';
update points set Opposition = 'Sheffield Wednesday' where Opposition = 'Sheffield Weds';
