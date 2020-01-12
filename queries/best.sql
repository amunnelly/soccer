select team, avg(goalsAgainst) as avg_goalsAgainst
from points
group by team
order by avg(goalsAgainst) desc;
