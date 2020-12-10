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
    let mpuTemperatureData = await fetch("/api/v0/packet/temperature/mpu?last=" + lastID);
    mpuTemperatureData = await mpuTemperatureData.json()
    await Plotly.newPlot('chart', [
        {
            y: data['bmpTemps'],
            type: 'line',
	        yaxis: {
		        autorange: true,
	    }
        }, {
            y: mpuTemperatureData['mpuTemps'],
            type: 'line',
            // yaxis: {
            //     autorange: true,
            // },
        }
    ]);
    count += data['bmpTemps'].length;


    let pressureChart = document.createElement("div");
    pressureChart.id = "pressureChart";
    document.getElementsByTagName('body')[0].appendChild(pressureChart);
    let pressureData = await fetch("/api/v0/packet/pressure?last=" + lastPressureID);
    pressureData = await pressureData.json();
    await Plotly.newPlot('pressureChart', [{
        // title: {
        //     text: 'Pressure in Pa',
        //     font: {
        //         family: 'Courier New, monospace',
        //         size: 24
        //     },
        // },
        y: pressureData['pressures'],
        type: 'line',
        yaxis: {
            autorange: true,
            // title: {
            //     text: 'Pressure in Pa',
            //     font: {
            //         family: 'Courier New, monospace',
            //         size: 18,
            //         color: '#7f7f7f'
            //     }
            // }
        },
        marker: {
            color: 'rgb(255, 0, 0)'
        },
    }]);
    pressureCount = pressureData['pressures'].length;

    setInterval(async () => {
        data = await fetch('/api/v0/packet/temperature/bmp?last=' + lastID);
        data = await data.json();
        mpuTemperatureData = await fetch('/api/v0/packet/temperature/mpu?last=' + lastID)
        mpuTemperatureData = await mpuTemperatureData.json();
	    await Plotly.extendTraces(['chart', 'chart'], { y: [data['bmpTemps'], mpuTemperatureData['bmpTemps']] }, [0, 1]);
        lastID = data['IDs'][data['IDs'].length - 1];
        count += data['bmpTemps'].length;

        if (count > 500) {
            Plotly.relayout('chart', {
                marker: {
                    color: 'rgb(255, 0, 0)'
                },
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
        await Plotly.extendTraces('pressureChart', { y: [pressureData['pressures']]}, [0]);
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
    }, 500);
})();
