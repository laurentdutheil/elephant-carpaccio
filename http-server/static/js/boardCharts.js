const currencyFormat = new Intl.NumberFormat('en-US', {style: 'currency', currency: 'USD', maximumFractionDigits: 0});
const tooltipLabelFunction = function (context) {
    let label = context.dataset.label || '';

    if (label) {
        label += ': ';
    }
    if (context.parsed.y !== null) {
        label += currencyFormat.format(context.parsed.y);
    }
    return label;
}

scoreChart = new Chart(document.getElementById('iterationScores'), {
    type: 'line',
    data: {
        labels: ['IT0', 'IT1', 'IT2', 'IT3', 'IT4', 'IT5', 'IT6'],
        datasets: [
            {
                label: 'ideal',
                data: [0, 3, 9, 14, 17, 18, 18],
                borderWidth: 1,
                tension: 0.3,
                borderDash: [10, 5]
            },
        ],
    },
    options: {
        aspectRatio: 1.2,
        plugins: {
            title: {
                display: true,
                text: ["STORY POINTS", "1 story = 1 point"],
                font: {
                    size: 14,
                    weight: 'normal'
                }
            },
            colors: {
                forceOverride: true
            }
        }
    }
});

businessValueChart = new Chart(document.getElementById('iterationBusinessValues'), {
    type: 'line',
    data: {
        labels: ['IT0', 'IT1', 'IT2', 'IT3', 'IT4', 'IT5', 'IT6'],
        datasets: [
            {
                label: 'ideal',
                data: [0.00, 1000.00, 12000.00, 13000.00, 13400.00, 13400.00, 13400.00],
                borderWidth: 1,
                tension: 0.3,
                borderDash: [10, 5]
            },
        ],
    },
    options: {
        aspectRatio: 1.2,
        scales: {
            y: {
                ticks: {
                    callback: function (value) {
                        return currencyFormat.format(value)
                    }
                }
            }
        },
        plugins: {
            title: {
                display: true,
                text: "BUSINESS VALUE",
                font: {
                    size: 14,
                    weight: 'normal'
                }
            },
            colors: {
                forceOverride: true
            },
            tooltip: {
                callbacks: {
                    label: tooltipLabelFunction
                }
            }
        }
    }
});

riskChart = new Chart(document.getElementById('risk'), {
    type: 'line',
    data: {
        labels: ['IT0', 'IT1', 'IT2', 'IT3', 'IT4', 'IT5', 'IT6'],
        datasets: [
            {
                label: 'ideal',
                data: [100, 50, 17, 4, 1, 0, 0],
                borderWidth: 1,
                tension: 0.3,
                borderDash: [10, 5]
            },
        ],
    },
    options: {
        aspectRatio: 1.2,
        plugins: {
            title: {
                display: true,
                text: "RISK",
                font: {
                    size: 14,
                    weight: 'normal'
                }
            },
            colors: {
                forceOverride: true
            }
        }
    }
});

costOfDelayChart = new Chart(document.getElementById('iterationCostOfDelay'), {
    type: 'line',
    data: {
        labels: ['IT0', 'IT1', 'IT2', 'IT3', 'IT4', 'IT5', 'IT6'],
        datasets: [
            {
                label: 'ideal',
                data: [0.00, 0.00, 0.00, 0.00, 0.00, 0.00, 0.00],
                borderWidth: 1,
                tension: 0.3,
                borderDash: [10, 5]
            },
        ],
    },
    options: {
        aspectRatio: 1.2,
        scales: {
            y: {
                ticks: {
                    callback: function (value) {
                        return currencyFormat.format(value)
                    }
                }
            }
        },
        plugins: {
            title: {
                display: true,
                text: ["COST OF DELAY", "business value lost"],
                font: {
                    size: 14,
                    weight: 'normal'
                }
            },
            colors: {
                forceOverride: true
            },
            tooltip: {
                callbacks: {
                    label: tooltipLabelFunction
                }
            }
        }
    }
});