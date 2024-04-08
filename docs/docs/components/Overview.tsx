import React from "react";
import "../styles.css";

export default function Overview() {
  return (
    <>
      <div className="voc_Div">
        <div className="flex max-w-3xl mx-auto flex-row">
          <div className="border-dotted border-border border p-5">
            <header className="vocs_Header border-none">
              <h1 className="vocs_H1 vocs_Heading">What is BeaconKit</h1>
            </header>
            <p className="vocs_Paragraph">
              BeaconKit introduces an innovative framework that utilizes the
              Cosmos-SDK to create a flexible, customizable consensus layer
              tailored for Ethereum-based blockchains
            </p>
          </div>
          <div className="border-dotted border-border border p-5">
            <header className="vocs_Header border-none">
              <h1 className="vocs_H1 vocs_Heading">Features</h1>
            </header>
            <p className="vocs_Paragraph">
              BeaconKit introduces an innovative framework that utilizes the
              Cosmos-SDK to create a flexible, customizable consensus layer
              tailored for Ethereum-based blockchains
            </p>
          </div>
        </div>
        <header className="vocs_Header border-dotted">
          <h1 className="vocs_H1 vocs_Heading">Sponsors</h1>
        </header>
      </div>
    </>
  );
}
