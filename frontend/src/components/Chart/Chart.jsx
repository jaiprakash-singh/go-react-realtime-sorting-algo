import React from "react";
import BarChart from 'react-bar-chart';

const margin = { top: 50, right: 50, bottom: 50, left: 100 };

const Chart = (props) => {
    let data = [];
    if(props.data !== undefined) {
        let propsData = props.data.data;
        if (propsData !== undefined) {
            console.log("PropsData:", propsData.split(","))
            data = propsData.split(",").map(function (s) {
                return { text: s.split(":")[0], value: s.split(":")[1] }
            });
        }
    }
    console.log("Data:", data)

    // useEffect(() => {
    //     window.onresize = () => {
    //         this.setState({ width: this.refs.root.offsetWidth });
    //     };
    // }, []);

    const handleBarClick = (element, id) => {
        console.log(`The bin ${element.text} with id ${id} was clicked`);
    }

    return (
        data !== [] &&
        <div style={{ width: '80%' }}>
            <div className="row">
                <BarChart ylabel=''
                    width={800}
                    height={400}
                    margin={margin}
                    data={data}
                    onBarClick={handleBarClick} />
            </div>
            <div className="row">
                {props.status === "init" && <button style={{width:"100px"}} onClick={props.sort}>Sort</button> }
                {props.status === "started" && <button style={{width:"100px"}} onClick={props.stop}>Stop</button>}
                {props.status === "stopped" && <button style={{width:"100px"}} onClick={props.resume}>Resume</button>}
                {props.status === "finished" && <button style={{width:"100px"}} onClick={props.newchart}>New Chart</button>}
            </div>
        </div>
    );
}

export default Chart;


    // const data = [
    //     { text: 'One', value: 500 },
    //     { text: 'Two', value: 350 },
    //     { text: 'Three', value: 250 },
    //     { text: 'Four', value: 450 },
    //     { text: 'Five', value: 720 },
    //     { text: 'Six', value: 25 },
    //     { text: 'Seven', value: 150 },
    // ];
