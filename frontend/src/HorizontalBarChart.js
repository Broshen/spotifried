import React from 'react';
import Chart from 'chart.js';
import Button from 'react-bootstrap/Button';


class HorizontalBarChart extends React.Component {
	constructor(props) {
		super(props);
		this.canvasRef = React.createRef();

		this.shown = Math.min(5, this.props.labels.length)
		this.labelsShown = this.props.labels.slice(0,this.shown)
		this.valuesShown = this.props.values.slice(0,this.shown)
	}

	componentDidUpdate() {
		this.chart.data.labels = this.props.labels.slice(0, this.shown)
		this.chart.data.datasets[0].data = this.props.values.slice(0,this.shown)
		this.chart.update();
	}

	componentDidMount() {
		this.chart = new Chart(this.canvasRef.current, {
			type: 'horizontalBar',
			options: {
				legend: {
					display: false,
				},
				scales: {
					yAxes: [{
						ticks: {
							fontColor: 'white'
						},
					}],
					xAxes: [{
						ticks: {
							beginAtZero:true,
							fontColor: 'white'
						},
					}]
				},
				maintainAspectRatio: true,
				aspectRatio: 1,
				tooltips: {
			      callbacks: {
			        label: function(tooltipItem, data) {
			        	return data.labels[tooltipItem.index] + " : " + data.datasets[0].data[tooltipItem.index] + " " + this.props.tooltipUnit;
				        }.bind(this)
				    },
				},
			},
			data: {
				labels: this.labelsShown,
				datasets: [{
					data: this.valuesShown,
				backgroundColor: "rgba(255,255,255,1)",
				}]
			}
		});

	}

	loadMore() {
		this.shown = Math.min(this.shown + 1, this.props.labels.length)
		this.chart.data.labels = this.props.labels.slice(0, this.shown)
		this.chart.data.datasets[0].data = this.props.values.slice(0,this.shown)
		this.chart.update();
	}

	render() {
		return (
			<div>
				<canvas ref={this.canvasRef} />
				{this.shown < this.props.labels.length && <Button onClick={this.loadMore.bind(this)}>Load More</Button>}
			</div>
		)
	}

}

export default HorizontalBarChart


