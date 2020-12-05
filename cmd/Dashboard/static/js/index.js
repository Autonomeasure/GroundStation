let lastID = 0;
let count = 0;

(async () => {
    let data = await fetch('/api/v0/packet/temperature/bmp?last=' + lastID);
    data = await data.json();
    Plotly.plot('chart', [{
        y: data['bmpTemps'],
        type: 'line',
    }]);

    setInterval(async () => {
        data = await fetch('/api/v0/packet/temperature/bmp?last=' + lastID);
        data = await data.json();
        Plotly.extendTraces('chart', { y: [data['bmpTemps']]}, [data['bmpTemps'].length - 1]);
        count++;

        if (count > 500) {
            Plotly.relayout('chart', {
                xaxis: {
                    range: [count - 500, count]
                }
            });
        }
    }, 200)
})();