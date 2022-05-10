import React from "react";
import "./InteractiveChart.css";
import {BarChartContainer, Number, MakeBar} from "./style";

const InteractiveChart = props => {

    const __DATA__ = [
        { distance: 13, colors: ["#ffd847", "#e0a106"] },
        { distance: 20, colors: ["#ff47ab", "#e0064e"] },
        { distance: 16, colors: ["#add9c0", "#1da890"] },
        { distance: 30, colors: ["#cbd9ad", "#7ca81d"] },
        { distance: 22, colors: ["#d9c1ad", "#714511"] },
    ];

    const Chart = __DATA__.map(({ distance, colors }, i) => {
        return (
            <BarChartContainer key={i}>
                <Number color={colors[1]}>{distance} km</Number>
                <MakeBar height={distance * 2} colors={colors} />
            </BarChartContainer>
        );
    })

    return Chart;

}

export default InteractiveChart;