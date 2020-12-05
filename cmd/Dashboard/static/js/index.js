let lastID = 0;
let count = 0;
let length = 0;

(async () => {
    let chart = document.createElement("div");
    chart.id = "chart";
    document.getElementsByTagName('body')[0].appendChild(chart);
    let data = await fetch('/api/v0/packet/temperature/bmp?last=' + lastID);
    data = await data.json();
    Plotly.plot('chart', [{
        y: data['bmpTemps'],
        type: 'line',
	yaxis: {
		autorange: true,
	}
    }]);
    count += data['bmpTemps'].length;

    setInterval(async () => {
        data = await fetch('/api/v0/packet/temperature/bmp?last=' + lastID);
        data = await data.json();
        Plotly.extendTraces('chart', { y: [data['bmpTemps']]}, [0]);
        // count++;
        lastID = data['IDs'][data['IDs'].length - 1];
        count += data['bmpTemps'].length;

        if (count > 500) {
            Plotly.relayout('chart', {
                xaxis: {
                    range: [count - 500, count],
                },
		yaxis: {
			autorange: true,
		}
            });
        }
    }, 200)
})();
