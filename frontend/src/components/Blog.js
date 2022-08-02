import React, {useState} from 'react';
/* import axios from 'axios';
import DeleteIcon from '@material-ui/icons/Delete';
import Button from '@material-ui/core/Button';
import EditIcon from '@mui/icons-material/Edit'; */
import TextField from '@material-ui/core/TextField';
import Api from './Api';

function Blog(props) {
  const {blog} = props;

  const [title, setTitle] = useState('')
    const [description, setDescription] = useState('')
    
    const handleTitle = (e) => {
        setTitle(e.target.value)
        console.log(setTitle)
        console.log(title)
    }

    const handleDescription = (e) => {
        setDescription(e.target.value)
        console.log(setDescription)
        console.log(description)
    }

  /* const deletes = async (id) => {
  axios.delete(`http://localhost:3001/blogs/${id}`, {
  }).then(res => {
      console.log(res)
  });
  }

  const update = async (id) => {
    axios.put(`http://localhost:3001/blogs/${id}`, {
      title:title,
      description:description
    }).then(res => {
      console.log(res)
      console.log(title,description)
    });
  } */

    return(
      <>
      { blog &&
        <div style={{maxWidth:"200px",minWidth:"200px", height:"200px", border:"solid 1px",borderRadius:"4px", margin:"4px",display:"block",justifyContent:"center"}}>
          {/* <div>
              <Button onClick={()=>deletes(blog && blog.id)}>
              <DeleteIcon/>
              </Button>
              <Button onClick={()=>update(blog && blog.id)}>
              <EditIcon/>
              </Button>
          </div> */}
          <Api></Api>
          <div>
            <TextField
              id="title"
              label="Title"
              variant="outlined"
              color="secondary"
              value={title}
              onChange={handleTitle}
            />
            <TextField
              id="description"
              label="Description"
              variant="outlined"
              color="secondary"
              value={description}
              onChange={handleDescription}
            />
          </div>
            {blog && blog.title}
            <div>
            {blog && blog.description}
            </div>
        </div>
        }
      </>
    )
}

export default Blog;