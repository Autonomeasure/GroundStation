let lastID;
let count;

const getData = async (lastID = 0) => {
    const res = await fetch('/api/v0/packet?last=' + lastID);
    const json = await res.json();
    // console.log(json);
    lastID = json[json.length - 1].id;
    return json;
}

(async () => {
    Plotly.plot('chart', [{
        y: getData(),
        type: 'line',
    }]);

    setInterval(() => {
        let data = getData(lastID);
        Plotly.extendTraces('chart', { y: [data]}, [data.length - 1]);
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