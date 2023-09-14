type Props = { title: string; content: string }

function Note(props: Props) {
  return (
    <div>
      Note
      <h2>{props.title}</h2>
      <p>{props.content}</p>
    </div>
  )
}

export default Note
