const es = new EventSource("/sse");
es.onerror = (err) => {
    console.log("onerror", err)
    es.close()
};

es.addEventListener("score", (event) => {
    const parsedData = JSON.parse(event.data);
    const {teamName, newScore, newBusinessValue, newRisk, newCostOfDelay} = parsedData;

    scoreChart.data.datasets
        .find((dataset) => dataset.label === teamName)
        .data.push(newScore)
    scoreChart.update()

    businessValueChart.data.datasets
        .find((dataset) => dataset.label === teamName)
        .data.push(newBusinessValue)
    businessValueChart.update()

    riskChart.data.datasets
        .find((dataset) => dataset.label === teamName)
        .data.push(newRisk)
    riskChart.update()

    costOfDelayChart.data.datasets
        .find((dataset) => dataset.label === teamName)
        .data.push(-newCostOfDelay)
    costOfDelayChart.update()
});

es.addEventListener("registration", (event) => {
    const parsedData = JSON.parse(event.data);
    const {teamName} = parsedData;

    scoreChart.data.datasets
        .push({label: teamName, data: [0], borderWidth: 1, tension: 0.3})
    scoreChart.update()

    businessValueChart.data.datasets
        .push({label: teamName, data: [0.00], borderWidth: 1, tension: 0.3})
    businessValueChart.update()

    riskChart.data.datasets
        .push({label: teamName, data: [100], borderWidth: 1, tension: 0.3})
    riskChart.update()

    costOfDelayChart.data.datasets
        .push({label: teamName, data: [0.00], borderWidth: 1, tension: 0.3})
    costOfDelayChart.update()
});
