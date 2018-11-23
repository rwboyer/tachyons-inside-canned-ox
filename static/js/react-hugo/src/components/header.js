import React from "react"

export default () => (
  <header className='tc pv4 pv5-ns'>
    <h1 className='f5 f4-ns fw6 mid-gray'>Jasper Whitehouse</h1>
    <h2 className='f6 gray fw2 ttu tracked'>Los Angeles</h2>
  </header>
)

export const List = ({ title, children }) => (
  <article className='pa3 pa5-ns'>
    <h1 className='f4 bold center mw6' children={title} />
    <ul className='list pl0 ml0 center mw6 ba b--light-silver br2' children={children} />
  </article>
)

export const ListItem = ({ children }) => (
  <li className='ph3 pv3 bb b--light-silver' children={children} />
)

export class TachLikeButton extends React.Component {
  constructor(props) {
    super(props);
    this.state = { liked: false };
    this.handleClick = this.handleClick.bind(this);
  }

  handleClick(e) {
    e.preventDefault();
    this.setState(state => ({
      liked: !state.liked
    }));
  }


  render() {
//    if (this.state.liked) {
//      return 'You liked this.';
//    }

//    return e(
//      'button',
//      { onClick: () => this.setState({ liked: ! this.state.liked }) },
//      'Like'
//    );
  return <a href="#" onClick = { this.handleClick } className="dib mt4 link w-auto-ns bg-transparent dark-red f5 pv2 pv3-ns ph4 ba b--dark-red dim ttu">{this.state.liked ? 'Liked' : 'Like'}</a>

  }

}
