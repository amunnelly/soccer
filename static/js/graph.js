function drawGraph(files) {

    var data = files[0];
    var colors = files[1];

    const width = 800
    const height = 493

    const margin = {
        left: 50,
        right: 10,
        top: 10,
        bottom: 50
    }

    var plot_width = width - margin.left - margin.right;
    var plot_height = height - margin.top - margin.bottom;


    teams = d3.nest()
        .key(d => { return d.Team })
        .rollup(e => {
            return {
                gd: d3.sum(e, d => { return d.GoalDifference }),
                points: d3.sum(e, d => { return d.Points })
            }
        })
        .entries(data)

    let maxGD = 0;
    let minGD = 0;

    teams.forEach(d => {
        if (d.value.gd > maxGD) {
            maxGD = d.value.gd
        }
        if (d.value.gd < minGD) {
            minGD = d.value.gd;
        }
    })


    var scaleX = d3.scaleLinear()
        .domain([0, 38 * 3])
        .range([0, plot_width])

    var scaleY = d3.scaleLinear()
        .domain([minGD, maxGD])
        .range([plot_height, 0])

    var canvas = d3.select('#canvas')
        .append('svg')
        .style('background', '#b5fcb5')
        .attr('height', height)
        .attr('width', width)

    var plot = canvas.append('g')
        .attr("transform", "translate(" + margin.left + "," + margin.top + ")")


    var dots = plot.selectAll('circle').data(teams).enter()
        .append('circle')
        .attr('class', 'point-circle')
        .attr("cx", d => { return scaleX(d.value.points) })
        .attr("cy", d => { return scaleY(d.value.gd) })
        .attr("r", 10)
        .style("stroke-width", 2)
        .style("fill", d => { return colors[d.key].first })
        .style("stroke", d => { return colors[d.key].second })

    var tooltip = d3.select("body")
        .append('g')
        .attr('class', 'tooltip')
        .style('opacity', 0)


    dots.on("mouseover", function (d) {
        d3.select(event.currentTarget)
            .attr("r", 15);

        tooltip
            .style('opacity', 0.9)
            .style("left", (d3.event.pageX + 10) + "px")
            .style("top", (d3.event.pageY - 10) + "px")
            .html(tooltipFormatter(d))
    })
        .on("mouseout", function (d) {
            d3.select(event.currentTarget)
                .attr("r", 10);

            tooltip
                .style('opacity', 0)
        })


    var xAx = d3.axisBottom(scaleX);
    var yAx = d3.axisLeft(scaleY);

    plot.append('g')
        .attr("transform",
            "translate(0, " + scaleY(0) + ")")
        .attr('class', 'axes')
        .call(xAx);

    plot.append('g')
        .attr("transform",
            "translate(0, 0)")
        .attr('class', 'axes')
        .call(yAx);

    canvas.append('text')
        .attr("x", plot_width - 20)
        .attr("y", scaleY(0) + 40)
        .attr("class", "axis-label")
        .text('Points')

    canvas.append("text")
        .attr("transform", "rotate(-90)")
        .attr("y", 0)
        .attr("x", 0 - (plot_height / 2))
        .attr("dy", "1em") // nudge the text; see p 65 of Tips and Tricks
        .attr("class", "axis-label")
        .style("text-anchor", "middle")
        .text("Goal Difference");



    function tooltipFormatter(d) {
        var team = "<strong>" + d.key + "</strong></br>"
        var gd = "GD: " + d.value.gd
        var points = "Points: " + d.value.points + "</br>"
        return team + points + gd

    }

}


function draw() {
    Promise.all([d3.csv("./static/data/data.csv"),
    d3.json("./static/colors.json")])
        .then(function (files) {
            files[0].forEach(d => {
                d.Points = +d.Points
                d.GoalDifference = +d.GD
            })

            drawGraph(files)
        })
        .catch((err) => {
            console.log("Hell's bells, something's gone wrong.")
            console.log(err)
        })
}