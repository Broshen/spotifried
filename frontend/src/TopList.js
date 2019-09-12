import React from 'react';

class TopList extends React.Component {
	render() {
		return (
			<div className="list-group fixed-list">
				{this.props.elements}
			</div>
		)
	}

}

export default TopList


