function displaySynonymsChart(synonyms, similarities) {
    const chartElement = document.getElementById("synonyms-chart");
    const chartData = {
        labels: synonyms,
        datasets: [
            {
                label: "Similarity",
                data: similarities,
                backgroundColor: "rgba(75, 192, 192, 0.2)",
                borderColor: "rgba(75, 192, 192, 1)",
                borderWidth: 1,
            },
        ],
    };
    const chartOptions = {
        plugins: {
            legend: {
                display: false,
            },
            tooltip: {
                enabled: true,
                callbacks: {
                    label: (context) => {
                        const label = context.label || "";
                        const value = context.raw;
                        return `${label}: Similarity ${value}`;
                    },
                },
            },
        },
        scales: {
            x: {
                grid: {
                    display: false,
                },
            },
            y: {
                min: 0,
                max: 10,
                grid: {
                    display: false,
                },
                ticks: {
                    display: true,
                },
            },
        },
    };

    if (window.synonymsChart) {
        window.synonymsChart.destroy();
    }

    window.synonymsChart = new Chart(chartElement, {
        type: "bar",
        data: chartData,
        options: chartOptions,
    });
}

// Update the event listener to pass the similarities data
document.getElementById("search-form").addEventListener("submit", async (e) => {
    e.preventDefault();

    const searchInput = document.getElementById("search-input");
    const word = searchInput.value.trim();

    if (!word) {
        alert("Please enter a word");
        return;
    }

    try {
        const response = await fetch(`/api/synonyms?word=${word}`);
        if (!response.ok) {
            throw new Error(await response.text());
        }
        const data = await response.json();
        displaySynonymsChart(data.synonyms, data.similarities);
    } catch (error) {
        alert(`Error: ${error.message}`);
    }
});
