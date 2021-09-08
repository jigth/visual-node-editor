import axios from 'axios';

export async function getURL(url) {
    try {
        const { data } = await axios.get(url);
        console.log({ data });
        return data;
    } catch (err) {
        console.error(err);
    }
}

export async function postURL(url, reqBody) {
    try {
        const { data } = await axios.post(url, reqBody);
        console.log({ data });
        return data;
    } catch (err) {
        console.error(err);
    }
}

// Use this function to do a quick check if the backend and frontend communicate
// well
export async function requestBackend() {
    const post = {
        CodeID: '123',
        Name: 'hehe.jpg',
        Code: 'Lorem ipsum gatorum madonorum',
    };

    await getURL('http://localhost:3000/');
    await postURL('http://localhost:3000/code/save', post);
}
