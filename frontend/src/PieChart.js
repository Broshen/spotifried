import React from 'react';
import Chart from 'chart.js';


class PieChart extends React.Component {
	constructor(props) {
		super(props);
		this.canvasRef = React.createRef();
	}

	componentDidUpdate() {
		this.chart.data.labels = this.props.labels;
		this.chart.data.datasets[0].data = this.props.values;
		this.chart.update();
	}

	componentDidMount() {
		this.chart = new Chart(this.canvasRef.current, {
			type: 'pie',
			options: {
				legend: {
					display: false,
				},
				maintainAspectRatio: true,
				aspectRatio: 1,
				tooltips: {
			      callbacks: {
			        label: function(tooltipItem, data) {
			        	return data.labels[tooltipItem.index] + " : " + parseInt(100*data.datasets[0].data[tooltipItem.index]) + "%"
				        }
				    },
				},
			},
			data: {
				labels: this.props.labels,
				datasets: [{
					data: this.props.values,
				}]
			}
		});

	}

	render() {
		return <canvas ref={this.canvasRef} />;
	}

}

export default PieChart


