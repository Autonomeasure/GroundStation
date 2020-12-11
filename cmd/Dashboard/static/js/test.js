const getBMPData = async (last = 0) => {
    return await (await fetch("/api/v0/packet/temperature/bmp")).json()
}

const getMPUData = async (last = 0) => {
    return await (await fetch("/api/v0/packet/temperature/mpu")).json()
}

let lastTemperatureID;
let temperatureCount;
let bmpTempData, mpuTempData;

let temperatureChart;

(async () => {
    temperatureChart = document.createElement("div");
    document.getElementsByTagName("body")[0].appendChild(temperatureChart);

    bmpTempData = await getBMPData();
    mpuTempData = await getMPUData();

    Plotly.newPlot(temperatureChart, [
        {
            y: bmpTempData['bmpTemps'],
            type: 'line',
            name: 'BMP280 temperature',
            yaxis: {
                autorange: true,
            }
        }, {
            y: mpuTempData['mpuTemps'],
            type: 'line',
            name: 'MPU6050 temperature',
            yaxis: {
                autorange: true,
            }
        }
    ], {
        showlegend: true,
        legend: {
            x: 1,
            xanchor: 'right',
            y: 1
        }
    });

    temperatureCount += bmpTempData['bmpTemps'].length;

    setInterval(async () => {
        bmpTempData = await getBMPData();
        mpuTempData = await getMPUData();

        Plotly.extendTraces([temperatureChart, temperatureChart], { y: [bmpTempData['bmpTemps'], mpuTempData['mpuTemps']] }, [0, 1]);
        lastTemperatureID = bmpTempData['IDs'][bmpTempData['IDs'].length - 1];
        temperatureCount += bmpTempData['bmpTemps'].length;

        if (temperatureCount > 500) {
            Plotly.relayout(temperatureChart, {
               marker: {
                   color: 'rgb(255, 0, 0)'
               }, xaxis: {
                   range: [temperatureCount - 500, temperatureCount],
                }, yaxis: {
                   autorange: true,
                }
            });
        }
    }, 500);
})();