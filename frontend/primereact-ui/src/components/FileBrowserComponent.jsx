import React, { useEffect, useState } from 'react';
import { Tree } from 'primereact/tree';
import { useImperativeHandle, forwardRef } from 'react';

const FileBrowserComponent = forwardRef((props, ref) => {
  useImperativeHandle(ref, () => ({
    childFunction() {
      fetch('/files2')
          .then(response => response.json())
          .then(data => {
            var x = transformFileTree(data);
            setFileTree(x);
          });

      fetch('/files3')
          .then(response => response.json())
          .then(data => {
            setFileTree2(data);
          });
    }
  }));
  const [fileTree, setFileTree] = useState(null);
  const [fileTree2, setFileTree2] = useState(null);
  const [selectedKey, setSelectedKey] = useState('');

  useEffect(() => {
    fetch('/files2')
        .then(response => response.json())
        .then(data => {
          var x = transformFileTree(data);
          setFileTree(x);
        });

    fetch('/files3')
        .then(response => response.json())
        .then(data => {
          setFileTree2(data);
        });


  }, []);

  let keyCounter = 0;

  const generateKey = () => {
    return `${keyCounter++}`;
  };


   const findElementByKey = (key, nodes) => {
    for (const node of nodes) {

        if (node.key === key) {
            console.log('Found node with key:', key);
            return node;
        }

        if (node.children) {
            const result = findElementByKey(key, node.children);
            if (result) {
                return result;
            }
        }
    }
    return null;
};

    const FlatSearch = (key, stuff) => {
        for (let i = 0, len = stuff.length; i < len; i++) {
            if (i.toString() == key) {
                return stuff[i];
            }
        }
    }

  const transformFileTree = (file) => {
    const transformFile = (file) => {
      const key = generateKey();

      if (!file.children) {
        return {
          key: key,
          label: file.name,
          data: file.name + (file.isDir ? ' Folder' : ''),
          icon: file.isDir ? 'pi pi-fw pi-folder' : 'pi pi-fw pi-file',
          path: file.path,
        };
      } else {
        return {
          key: key,
          label: file.name,
          data: file.name + ' Folder',
          icon: file.isDir ? 'pi pi-fw pi-folder' : 'pi pi-fw pi-file',
          path: file.path,
          children: file.children.map(child => transformFile(child)),
        };
      }
    };

    return file.children.map(child => transformFile(child));
  };

  const handleSelectionChange = (event) => {
    setSelectedKey(event.value);
    //console.log(fileTree2)
    //console.log(fileTree)

    var z = FlatSearch(event.value, fileTree2)
    //console.log(z)
    fetch('/send-data', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(z),
    })
        .then(response => response.text())
        .then(data => {
        })
        .catch(error => {
          console.error('Error:', error);
        });
    const element = findElementByKey('5', fileTree);
    console.log(element);
    //if (element) {
      //console.log(element.label)
    //}
  };
  


  return (
      <div>
        {fileTree && <Tree value={fileTree} selectionMode="single" selectionKeys={selectedKey} onSelectionChange={handleSelectionChange} className="w-full md:w-30rem" />}
      </div>
  );
})

export default FileBrowserComponent;
