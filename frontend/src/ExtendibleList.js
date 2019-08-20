import React from 'react';
import Button from 'react-bootstrap/Button';

export function ArtistListItem(props){
	return (
		<div className="list-group-item text-right" key={props.Id}>
			{props.Name}
		</div>
	)
}

export function SongListItem(props){
	// var artists = ""

	// if (props.Track.Artists.length == 1){
	// 	artists = props.Track.Artists[0].Name
	// }

	return (
		<div className="list-group-item text-left" key={props.Track.Id}>
			{props.Track.Name} by {props.Track.Artists.map(e => e.Name).join(", ")}
		</div>
	)
}

class ExtendibleList extends React.Component {
	constructor(props) {
		super(props);

		this.state = {
			shown: Math.min(5, this.props.data.length)
		}
	}

	// componentDidUpdate() {
	// 	this.chart.data.labels = this.props.labels.slice(0, this.shown)
	// 	this.chart.data.datasets[0].data = this.props.values.slice(0,this.shown)
	// 	this.chart.update();
	// }

	loadMore() {
		this.setState({shown: Math.min(this.state.shown + 5, this.props.data.length)})
		// this.chart.data.labels = this.props.labels.slice(0, this.shown)
		// this.chart.data.datasets[0].data = this.props.values.slice(0,this.shown)
		// this.chart.update();
	}

	render() {
		var rows = [];
		for(var i=0; i<this.state.shown; i++){
			rows.push(
				this.props.child(this.props.data[i])
			)
		}
		return (
			<div className="list-group">
				{rows}
				{this.state.shown < this.props.data.length && <Button onClick={this.loadMore.bind(this)}>Load More</Button>}
			</div>
		)
	}

}

export default ExtendibleList


