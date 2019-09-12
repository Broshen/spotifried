import React from 'react';

import Container from 'react-bootstrap/Container';
import Row from 'react-bootstrap/Row';
import Col from 'react-bootstrap/Col';
import Tooltip from 'react-bootstrap/Tooltip';
import OverlayTrigger from 'react-bootstrap/OverlayTrigger';
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';


import { FontAwesomeIcon } from '@fortawesome/react-fontawesome'

import PieChart from './PieChart';
import HorizontalBarChart from './HorizontalBarChart'
import TopList from './TopList'
import {api_url, frontend_url} from "./variables"

class Profile extends React.Component {
	constructor(props) {
		super(props);

		this.state = {
			data: null,
			error: null,
		};
	}

	componentDidMount() {
		fetch(api_url + 'profile/' + this.props.match.params.id)
		.then(response => response.json())
		.then(data => {
			this.setState({ data })
		})
		.catch(error => {
			this.setState({error})
		});
	}

	render () {

		var totalSongs = 0
		var totalGenres = 0
		if (this.state.error != null) {
			return (
				<Container className="center-page text-center">
				<h1>An error occurred. Try refreshing this page or reloading your spotify data</h1>
				<OverlayTrigger
				key="refresh"
				placement="bottom"
				overlay={
					<Tooltip id="tooltip-refresh">
					Re-fetch your data from Spotify to get updates since the last time you've refreshed (songs added or removed)
					</Tooltip>
				}
				>
				<Button variant="secondary" href={api_url + "fetch/" + this.props.match.params.id + "?nocache=" + Math.random()}> Reload
				</Button>
				</OverlayTrigger>
				</Container>)
		}

		else if(this.state.data != null){
			totalSongs = this.state.data.artists.reduce((a,b) => a + parseInt(b.SongCount),0)
			totalGenres = this.state.data.genres.reduce((a,b) => a + parseInt(b.ArtistCount),0)

			return(
				<Container>
				<Row className="pad-vertical">
				<Col xs={12} md={4}>
				Last Refreshed: {this.state.data.last_refreshed}	
				</Col>
				<Col xs={12} md={4} className="text-center">
				<h2>{this.state.data.username}</h2>
				</Col>
				<Col xs={12} md={4} className="text-right">
				<OverlayTrigger
				key="refresh"
				placement="bottom"
				overlay={
					<Tooltip id="tooltip-refresh">
					Re-fetch your data from Spotify to get updates since the last time you've refreshed (songs added or removed)
					</Tooltip>
				}
				>
				<Button variant="secondary" href={api_url + "fetch/" + this.props.match.params.id + "?nocache=" + Math.random()}> Reload
				</Button>
				</OverlayTrigger>
				</Col>
				</Row>
				
				<Row  className="pad-vertical">
				<Col xs={12} md={6}>
				<h3 className="text-center"> Your library, by genres: </h3>
				<PieChart
				labels={this.state.data.genres.map(d => d.Name)}
				values={this.state.data.genres.map(d => d.ArtistCount/totalGenres)}
				/>
				</Col>
				<Col xs={12} md={6} className="text-center">
				<h3> Your top genres: </h3>
				<HorizontalBarChart
				labels={this.state.data.genres.map(d => d.Name)}
				values={this.state.data.genres.map(d => d.ArtistCount)}
				tooltipUnit="artists"
				/>
				</Col>
				</Row>

				<Row  className="pad-vertical">
				<Col xs={12} md={6}>
				<h3 className="text-center"> Your library, by artists: </h3>
				<PieChart
				labels={this.state.data.artists.map(d => d.Name)}
				values={this.state.data.artists.map(d => d.SongCount/totalSongs)}
				/>
				</Col>
				<Col xs={12} md={6} className="text-center">
				<h3> Your top artists, by number of songs: </h3>
				<HorizontalBarChart
				labels={this.state.data.artists.map(d => d.Name)}
				values={this.state.data.artists.map(d => d.SongCount)}
				tooltipUnit="songs"
				/>
				</Col>
				</Row>

				<Row  className="pad-vertical">
				<Col xs={12}>
					<h3 className="text-center"> Your top artists, by affinity  &emsp;&emsp;&emsp;
						<OverlayTrigger
						key="affinity"
						placement="bottom"
						overlay={
							<Tooltip id="tooltip-affinity">
							Affinity is a measure of the expected preference a user has for a particular track or artist.  It is based on user behavior, including play history, but does not include actions made while in incognito mode. Light or infrequent users of Spotify may not have sufficient play history to generate a full affinity data set.
							</Tooltip>
						}
						>
							<FontAwesomeIcon icon="question-circle" />
						</OverlayTrigger>
					</h3>
				</Col>
				</Row>

				<Row className="pad-vertical hidden-xs">
				<Col xs={12} md={4} className="text-center">
					<h4> From the last 4 weeks </h4>
				</Col>
				<Col xs={12} md={4} className="text-center">
					<h4> Over the last 6 months </h4>
				</Col>
				<Col xs={12} md={4} className="text-center">
					<h4> Of all time</h4>
				</Col>
				</Row>

				<Row className="pad-vertical">
				<Col xs={12} md={4} className="text-center">
					<h4 className="show-xs-only pad-vertical"> From the last 4 weeks </h4>
					<TopList
						elements={this.state.data.top_artists[0].map((x, i) => (<div className="list-group-item text-left" key={i}>{(i+1) + ". " + x.Name}</div>))}
					/>
				</Col>
				<Col xs={12} md={4} className="text-center">
					<h4 className="show-xs-only pad-vertical"> Over the last 6 months </h4>
					<TopList
						elements={this.state.data.top_artists[1].map((x, i) => (<div className="list-group-item text-left" key={i}>{(i+1) + ". " + x.Name}</div>))}
					/>
				</Col>
				<Col xs={12} md={4} className="text-center">
					<h4 className="show-xs-only pad-vertical"> Of all time</h4>
					<TopList
						elements={this.state.data.top_artists[2].map((x, i) => (<div className="list-group-item text-left" key={i}>{(i+1) + ". " + x.Name}</div>))}
					/>
				</Col>
				</Row>



				<Row  className="pad-vertical">
				<Col xs={12}>
					<h3 className="text-center"> Your top tracks, by affinity  &emsp;&emsp;&emsp;
						<OverlayTrigger
						key="affinity"
						placement="bottom"
						overlay={
							<Tooltip id="tooltip-affinity">
							Affinity is a measure of the expected preference a user has for a particular track or artist.  It is based on user behavior, including play history, but does not include actions made while in incognito mode. Light or infrequent users of Spotify may not have sufficient play history to generate a full affinity data set.
							</Tooltip>
						}
						>
							<FontAwesomeIcon icon="question-circle" />
						</OverlayTrigger>
					</h3>
				</Col>
				</Row>

				<Row className="pad-vertical hidden-xs">
				<Col xs={12} md={4} className="text-center">
					<h4> From the last 4 weeks </h4>
				</Col>
				<Col xs={12} md={4} className="text-center">
					<h4> Over the last 6 months </h4>
				</Col>
				<Col xs={12} md={4} className="text-center">
					<h4> Of all time</h4>
				</Col>
				</Row>

				<Row className="pad-vertical">
				<Col xs={12} md={4} className="text-center">
					<h4 className="show-xs-only pad-vertical"> From the last 4 weeks </h4>
					<TopList
						elements={this.state.data.top_songs[0].map((x, i) => (<div className="list-group-item text-left" key={i}>{(i+1) + ". " + x.Name + " by " + x.Artists.map(e => e.Name).join(", ")}</div>))}
					/>
				</Col>
				<Col xs={12} md={4} className="text-center">
					<h4 className="show-xs-only pad-vertical"> Over the last 6 months </h4>
					<TopList
						elements={this.state.data.top_songs[1].map((x, i) => (<div className="list-group-item text-left" key={i}>{(i+1) + ". " + x.Name + " by " + x.Artists.map(e => e.Name).join(", ")}</div>))}
					/>
				</Col>
				<Col xs={12} md={4} className="text-center">
					<h4 className="show-xs-only pad-vertical"> Of all time</h4>
					<TopList
						elements={this.state.data.top_songs[2].map((x, i) => (<div className="list-group-item text-left" key={i}>{(i+1) + ". " + x.Name + " by " + x.Artists.map(e => e.Name).join(", ")}</div>))}
					/>
				</Col>
				</Row>





				<Row  className="pad-vertical">
				<Col xs={{span: 8, offset: 2}} className="text-center">
				<h4> Want to compare your tastes with a friend's? Share this link with them! </h4>
				<Form.Control type="text" value={frontend_url+"share/"+this.props.match.params.id} readOnly />
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

export default Profile