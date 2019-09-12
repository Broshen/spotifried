import React from 'react';
import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';

import TopList from './TopList';
import {api_url} from "./variables"

class Compare extends React.Component {
	constructor(props) {
		super(props);

		this.state = {
			data: null,
			error: null,
		};
	}

	componentDidMount() {
		fetch(api_url + 'compare/' + this.props.match.params.id1 + "/" + this.props.match.params.id2)
		.then(response => response.json())
		.then(data => {
			this.setState({ data })
		})
		.catch(error => {
			this.setState({error})
		});
	}

	render () {
		if (this.state.error != null) {
			return (
				<Container className="center-page text-center">
				<h1>An error occurred. Please try refreshing this page.</h1>
				</Container>)
		}

		else if(this.state.data != null){
			return(
				<Container>
					<Row className="pad-vertical"/>

					<Row className="text-center">
						<Col xs={12} md={3}>
							<h4><a href={"/profile/" + this.state.data.user1.id}>{this.state.data.user1.name}</a></h4>
						</Col>
						<Col xs={12} md={6} >
							<h5> and </h5>
						</Col>
						<Col xs={12} md={3}>
							<h4><a href={"/profile/" + this.state.data.user2.id}>{this.state.data.user2.name}</a></h4>
						</Col>
					</Row>
					<Row>
						<Col xs={12} md={{span: 8, offset:2}} className="text-center">
							have {this.state.data.songs.length} songs and {this.state.data.artists.length} artists in common!
							<br/>
							This is {Math.round(this.state.data.songs.length*20000/(this.state.data.user1.songcount + this.state.data.user2.songcount))/100}% of your combined libraries.
						</Col>
					</Row>

					<Row className="text-center pad-vertical">
						<Col xs={12} md={6} >
							<h4 className="pad-vertical"> Songs you've both saved: </h4>
							<TopList
								elements={
									this.state.data.songs.map(
										(x, i) => (
											<div className="list-group-item text-left" key={i}>
												{x.Track.Name + " by " + x.Track.Artists.map(e => e.Name).join(", ")}
											</div>)
								)}
							/>
						</Col>

						<Col xs={12} md={6} >
							<h4 className="pad-vertical"> Artists you both like: </h4>
							<TopList
								elements={
									this.state.data.artists.map(
										(x, i) => (
											<div className="list-group-item text-right" key={i}>
												{x.Name} - {x.SongCount} common songs
											</div>)
								)}
							/>
						</Col>
					</Row>



					<Row className="text-center pad-vertical">
						<Col xs={12} md={6} >
							<h4 className="pad-vertical"> Songs you both listen to a lot: </h4>
							<TopList
								elements={
									this.state.data.top_tracks.map(
										(x, i) => (
											<div className="list-group-item text-left" key={i}>
												{x.Name + " by " + x.Artists.map(e => e.Name).join(", ")}
											</div>)
								)}
							/>
						</Col>

						<Col xs={12} md={6} >
							<h4 className="pad-vertical"> Artists you both like a lot: </h4>
							<TopList
								elements={
									this.state.data.top_artists.map(
										(x, i) => (
											<div className="list-group-item text-right" key={i}>
												{x.Name}
											</div>)
								)}
							/>
						</Col>
					</Row>
				</Container>)
		}
		else {
			return(
				<Container className="center-page text-center">
				<h1>Loading...</h1>
				</Container>)
		}		
	}
}

export default Compare