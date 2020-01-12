select date,
hometeam,
FTHG,
FTAG,
awayteam from results
where awayteam = 'Leicester'
limit 1;