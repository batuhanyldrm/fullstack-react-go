import React from 'react'
import axios from 'axios';
import DeleteIcon from '@material-ui/icons/Delete';
import Button from '@material-ui/core/Button';
import EditIcon from '@mui/icons-material/Edit';

    const update = async (id,title,description) => {
        axios.put(`http://localhost:3001/blogs/${id}`, {
          title,
          description,
        }).then(res => {
          console.log(res)
          console.log(title,description)
        });
      }


    const deletes = async (id) => {
      axios.delete(`http://localhost:3001/blogs/${id}`, {
      }).then(res => {
          console.log(res)
      });
    }

function Api(blog) {
  return (
    <div>
      <Button onClick={()=>deletes(blog && blog.id)}>
      <DeleteIcon/>
      </Button>
      <Button onClick={()=>update(blog && blog.id)}>
      <EditIcon/>
      </Button>
    </div>
  )
}

export default Api