let lastID = 0;
let count = 0;
let lastPressureID = 0;
let pressureCount = 0;

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


    let pressureChart = document.createElement("div");
    chart.id = "pressureChart";
    document.getElementsByTagName('body')[0].appendChild(pressureChart);
    let pressureData = await fetch("/api/v0/packet/pressure?last=" + lastPressureID);
    pressureData = await pressureData.json();
    Plotly.plot('pressureChart', [{
        y: pressureData['pressures'],
        type: 'line',
        yaxis: {
            autorange: true,
        }
    }]);
    pressureCount = pressureData['pressures'].length;

    setInterval(async () => {
        data = await fetch('/api/v0/packet/temperature/bmp?last=' + lastID);
        data = await data.json();
	Plotly.extendTraces('chart', { y: [data['bmpTemps']]}, [0]);
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

        pressureData = await fetch("/api/v0/packet/pressure?last=" + lastPressureID);
        pressureData = await pressureData.json();
        Plotly.extendTraces('pressureChart', { y: [pressureData['pressures']]}, [0]);
        lastPressureID = pressureData['IDs'][data['IDs'].length - 1];
        pressureCount += pressureData['pressures'].length;

        if (pressureCount > 500) {
            Plotly.relayout('pressureChart', {
                xaxis: {
                    range: [pressureCount - 500, pressureCount],
                },
                yaxis: {
                    autorange: true,
                }
            });
        }
    }, 200);
})();
