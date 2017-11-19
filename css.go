package lcars

import ()

var lcarscss string = `
.lcars-gray-bg {
  background-color: rgba(34,51,85,0.467) !important;
}
.lcars-white-bg {
  background-color: #fff !important;
}
.lcars-black-bg {
  background-color: #000 !important;
}
.lcars-pale-canary-bg {
  background-color: #ff9 !important;
}
.lcars-golden-tanoi-bg {
  background-color: #fc6 !important;
}
.lcars-neon-carrot-bg {
  background-color: #f93 !important;
}
.lcars-eggplant-bg {
  background-color: #646 !important;
}
.lcars-lilac-bg {
  background-color: #c9c !important;
}
.lcars-anakiwa-bg {
  background-color: #9cf !important;
}
.lcars-mariner-bg {
  background-color: #36c !important;
}
.lcars-bahama-blue-bg {
  background-color: #069 !important;
}
.lcars-blue-bell-bg {
  background-color: #99c !important;
}
.lcars-melrose-bg {
  background-color: #99f !important;
}
.lcars-hopbush-bg {
  background-color: #c69 !important;
}
.lcars-chestnut-rose-bg {
  background-color: #c66 !important;
}
.lcars-orange-peel-bg {
  background-color: #f96 !important;
}
.lcars-atomic-tangerine-bg {
  background-color: #f90 !important;
}
.lcars-danub-bg {
  background-color: #68c !important;
}
.lcars-indigo-bg {
  background-color: #45b !important;
}
.lcars-lavender-purple-bg {
  background-color: #97a !important;
}
.lcars-cosmic-bg {
  background-color: #746 !important;
}
.lcars-red-damask-bg {
  background-color: #d64 !important;
}
.lcars-medium-carmine-bg {
  background-color: #a53 !important;
}
.lcars-bourbon-bg {
  background-color: #b62 !important;
}
.lcars-sandy-brown-bg {
  background-color: #e95 !important;
}
.lcars-periwinkle-bg {
  background-color: #cdf !important;
}
.lcars-dodger-blue-bg {
  background-color: #59f !important;
}
.lcars-dodger-blue-alt-bg {
  background-color: #36f !important;
}
.lcars-blue-bg {
  background-color: #01e !important;
}
.lcars-navy-blue-bg {
  background-color: #008 !important;
}
.lcars-husk-bg {
  background-color: #ba5 !important;
}
.lcars-rust-bg {
  background-color: #b41 !important;
}
.lcars-tamarillo-bg {
  background-color: #821 !important;
}
.lcars-gray-border {
  border-color: rgba(34,51,85,0.467) !important;
}
.lcars-white-border {
  border-color: #fff !important;
}
.lcars-black-border {
  border-color: #000 !important;
}
.lcars-pale-canary-border {
  border-color: #ff9 !important;
}
.lcars-golden-tanoi-border {
  border-color: #fc6 !important;
}
.lcars-neon-carrot-border {
  border-color: #f93 !important;
}
.lcars-eggplant-border {
  border-color: #646 !important;
}
.lcars-lilac-border {
  border-color: #c9c !important;
}
.lcars-anakiwa-border {
  border-color: #9cf !important;
}
.lcars-mariner-border {
  border-color: #36c !important;
}
.lcars-bahama-blue-border {
  border-color: #069 !important;
}
.lcars-blue-bell-border {
  border-color: #99c !important;
}
.lcars-melrose-border {
  border-color: #99f !important;
}
.lcars-hopbush-border {
  border-color: #c69 !important;
}
.lcars-chestnut-rose-border {
  border-color: #c66 !important;
}
.lcars-orange-peel-border {
  border-color: #f96 !important;
}
.lcars-atomic-tangerine-border {
  border-color: #f90 !important;
}
.lcars-danub-border {
  border-color: #68c !important;
}
.lcars-indigo-border {
  border-color: #45b !important;
}
.lcars-lavender-purple-border {
  border-color: #97a !important;
}
.lcars-cosmic-border {
  border-color: #746 !important;
}
.lcars-red-damask-border {
  border-color: #d64 !important;
}
.lcars-medium-carmine-border {
  border-color: #a53 !important;
}
.lcars-bourbon-border {
  border-color: #b62 !important;
}
.lcars-sandy-brown-border {
  border-color: #e95 !important;
}
.lcars-periwinkle-border {
  border-color: #cdf !important;
}
.lcars-dodger-blue-border {
  border-color: #59f !important;
}
.lcars-dodger-blue-alt-border {
  border-color: #36f !important;
}
.lcars-blue-border {
  border-color: #01e !important;
}
.lcars-navy-blue-border {
  border-color: #008 !important;
}
.lcars-husk-border {
  border-color: #ba5 !important;
}
.lcars-rust-border {
  border-color: #b41 !important;
}
.lcars-tamarillo-border {
  border-color: #821 !important;
}
.lcars-gray-color {
  color: rgba(34,51,85,0.467) !important;
}
.lcars-white-color {
  color: #fff !important;
}
.lcars-black-color {
  color: #000 !important;
}
.lcars-pale-canary-color {
  color: #ff9 !important;
}
.lcars-golden-tanoi-color {
  color: #fc6 !important;
}
.lcars-neon-carrot-color {
  color: #f93 !important;
}
.lcars-eggplant-color {
  color: #646 !important;
}
.lcars-lilac-color {
  color: #c9c !important;
}
.lcars-anakiwa-color {
  color: #9cf !important;
}
.lcars-mariner-color {
  color: #36c !important;
}
.lcars-bahama-blue-color {
  color: #069 !important;
}
.lcars-blue-bell-color {
  color: #99c !important;
}
.lcars-melrose-color {
  color: #99f !important;
}
.lcars-hopbush-color {
  color: #c69 !important;
}
.lcars-chestnut-rose-color {
  color: #c66 !important;
}
.lcars-orange-peel-color {
  color: #f96 !important;
}
.lcars-atomic-tangerine-color {
  color: #f90 !important;
}
.lcars-danub-color {
  color: #68c !important;
}
.lcars-indigo-color {
  color: #45b !important;
}
.lcars-lavender-purple-color {
  color: #97a !important;
}
.lcars-cosmic-color {
  color: #746 !important;
}
.lcars-red-damask-color {
  color: #d64 !important;
}
.lcars-medium-carmine-color {
  color: #a53 !important;
}
.lcars-bourbon-color {
  color: #b62 !important;
}
.lcars-sandy-brown-color {
  color: #e95 !important;
}
.lcars-periwinkle-color {
  color: #cdf !important;
}
.lcars-dodger-blue-color {
  color: #59f !important;
}
.lcars-dodger-blue-alt-color {
  color: #36f !important;
}
.lcars-blue-color {
  color: #01e !important;
}
.lcars-navy-blue-color {
  color: #008 !important;
}
.lcars-husk-color {
  color: #ba5 !important;
}
.lcars-rust-color {
  color: #b41 !important;
}
.lcars-tamarillo-color {
  color: #821 !important;
}
html,
body {
  font-family: "arial", monospace;
  color: #fff;
}
h1 {
  font-size: 270%;
  font-weight: bold;
}
h2 {
  font-size: 240%;
  font-weight: bold;
}
h3 {
  font-size: 200%;
  font-weight: bold;
}
h4 {
  font-size: 180%;
  font-weight: bold;
}
h5 {
  font-size: 160%;
  font-weight: bold;
}
h6 {
  font-size: 140%;
  font-weight: bold;
}
.lcars-app-container {
  display: flex;
  width: calc(100% - 1rem);
  height: calc(100% - 1rem);
  overflow: hidden;
  margin: 0.5rem;
}
.lcars-app-container #left-menu {
  position: fixed;
  padding-top: 4.5rem;
  height: 100%;
  left: 0.5rem;
}
.lcars-app-container #header {
  position: fixed;
  background-color: #000;
  width: calc(100% - 1rem);
  top: 0;
  left: 0.5rem;
  padding-top: 0.5rem;
  margin-bottom: 0.25rem;
  z-index: 1;
}
.lcars-app-container #footer {
  position: fixed;
  background-color: #000;
  bottom: 0;
  left: 0.5rem;
  padding-bottom: 0.5rem;
  width: calc(100% - 1rem);
  z-index: 1;
}
.lcars-app-container #container {
  margin-top: 0.25rem;
  margin-left: 0.25rem;
  padding-left: 9.5rem;
  padding-top: 4.5rem;
  padding-bottom: 4.5rem;
  width: 100%;
  height: 100%;
  overflow: auto;
}
.lcars-row {
  display: inline-flex;
  flex-direction: row;
  width: 100%;
}
.lcars-row.fill {
  flex: 1;
}
.lcars-row.centered {
  justify-content: center;
}
.lcars-row.full-centered {
  justify-content: center;
  align-items: center;
}
.lcars-row.right {
  justify-content: flex-end;
}
.lcars-row.right-centered {
  justify-content: flex-end;
  align-items: center;
}
.lcars-row.right-bottom {
  justify-content: flex-end;
  align-items: flex-end;
}
.lcars-row :last-child {
  margin-right: 0;
}
.lcars-row > * {
  margin-right: 0.25rem;
}
.lcars-column {
  display: inline-flex;
  flex-direction: column;
}
.lcars-column.fill {
  flex: 1;
}
.lcars-column.centered {
  justify-content: center;
}
.lcars-column.full-centered {
  justify-content: center;
  align-items: center;
}
.lcars-column.centered-right {
  justify-content: center;
  align-items: flex-end;
}
.lcars-column.bottom {
  justify-content: flex-end;
}
.lcars-column.bottom-centered {
  justify-content: flex-end;
  align-items: center;
}
.lcars-column :last-child {
  margin-bottom: 0;
}
.lcars-column > * {
  margin-bottom: 0.25rem;
}
.lcars-column.left-menu {
  padding-right: 2rem;
}
.lcars-column.right-menu {
  padding-left: 2rem;
}
.lcars-column.start-space {
  margin-top: 0.25rem;
}
.lcars-elbow {
  position: relative;
  width: 9.5rem;
  min-width: 9.5rem;
  height: 4.5rem;
  min-height: 4.5rem;
  background: #fc6;
  margin: 0;
}
.lcars-elbow a {
  display: inline-block;
  position: absolute;
  color: #000;
  bottom: 0;
  padding: 0.25rem;
}
.lcars-elbow:after {
  content: '';
  display: block;
  position: absolute;
  width: 2rem;
  height: 3rem;
  background: #000;
}
.lcars-elbow.left-bottom {
  border-top-left-radius: 3.75rem;
}
.lcars-elbow.left-bottom a {
  right: 2.25rem;
}
.lcars-elbow.left-bottom:after {
  right: 0;
  top: 1.5rem;
  border-top-left-radius: 1.875rem;
}
.lcars-elbow.left-top {
  border-bottom-left-radius: 3.75rem;
}
.lcars-elbow.left-top a {
  top: 0;
  right: 2.25rem;
}
.lcars-elbow.left-top:after {
  right: 0;
  bottom: 1.5rem;
  border-bottom-left-radius: 3.75rem;
}
.lcars-elbow.right-bottom {
  border-top-right-radius: 3.75rem;
}
.lcars-elbow.right-bottom:after {
  top: 1.5rem;
  border-top-right-radius: 1.875rem;
}
.lcars-elbow.right-bottom a {
  left: 2.25rem;
}
.lcars-elbow.right-top {
  border-bottom-right-radius: 3.75rem;
}
.lcars-elbow.right-top a {
  display: inline-block;
  top: 0;
  left: 2.25rem;
}
.lcars-elbow.right-top:after {
  bottom: 1.5rem;
  border-bottom-right-radius: 1.875rem;
}
.lcars-bar {
  position: relative;
  color: #000;
  height: 100%;
  width: 100%;
  margin: 0;
  display: inline-block;
  background-color: #fc6;
}
.lcars-bar:after {
  content: '';
  display: block;
  position: absolute;
  background: #000;
}
.lcars-bar.spacer {
  margin: 0;
  padding: 0;
  background-color: transparent;
  width: 0.25rem;
  min-width: 0.25rem;
  height: 0.25rem;
  min-height: 0.25rem;
}
.lcars-bar.double-spacer {
  background-color: transparent;
  width: 0.5rem;
  min-width: 0.5rem;
  height: 0.5rem;
  min-height: 0.5rem;
}
.lcars-bar.left-space {
  margin-left: 0.25rem;
}
.lcars-bar.left-double-space {
  margin-left: 0.5rem;
}
.lcars-bar.right-space {
  margin-right: 0.25rem;
}
.lcars-bar.right-double-space {
  margin-right: 0.5rem;
}
.lcars-bar.horizontal {
  height: 1.5rem;
}
.lcars-bar.left-end {
  width: 1.5rem;
  max-width: 1.5rem;
  min-width: 1.5rem;
  border-top-left-radius: 0.75rem;
  border-bottom-left-radius: 0.75rem;
}
.lcars-bar.left-end.decorated:after {
  right: 0.25rem;
  width: 0.25rem;
  height: 100%;
  background-color: #000;
}
.lcars-bar.right-end {
  width: 1.5rem;
  max-width: 1.5rem;
  min-width: 1.5rem;
  border-top-right-radius: 0.75rem;
  border-bottom-right-radius: 0.75rem;
}
.lcars-bar.right-end.decorated:after {
  left: 0.25rem;
  width: 0.25rem;
  height: 100%;
  background-color: #000;
  background-color: #000;
}
.lcars-bar.fill {
  flex: 1;
}
.lcars-bar.bottom {
  align-self: flex-end;
}
.lcars-bar .lcars-title {
  color: #fff;
  background-color: #000;
  text-transform: uppercase;
  margin: 0;
  margin-left: 0.833333333333335rem;
  padding-left: 0.166666666666667rem;
  padding-right: 0.166666666666667rem;
  padding-bottom: 0.166666666666667rem;
  display: inline-block;
  font-size: 150%;
  height: 100%;
}
.lcars-bar .lcars-title.right {
  float: right;
  margin-right: 0.833333333333335rem;
}
.lcars-element {
  position: relative;
  font-weight: bold;
  color: #000;
  text-align: right;
  background: #fc6;
  height: 3rem;
  width: 7.5rem;
  box-sizing: border-box;
  padding-left: 0.75rem;
  padding-right: 0.75rem;
  display: inline-flex;
  flex-direction: row;
  justify-content: flex-end;
  align-items: flex-end;
}
.lcars-element * {
  margin: 0;
}
.lcars-element.left-rounded {
  border-top-left-radius: 1.5rem;
  border-bottom-left-radius: 1.5rem;
  padding-left: 1.5rem;
}
.lcars-element.right-rounded {
  border-top-right-radius: 1.5rem;
  border-bottom-right-radius: 1.5rem;
  padding-right: 1.5rem;
}
.lcars-element.rounded {
  border-radius: 1.5rem;
  padding-left: 1.5rem;
  padding-right: 1.5rem;
}
.lcars-element.button {
  cursor: pointer;
}
.lcars-element.button:active {
  background-color: #c69;
}
.lcars-element .lcars-element-addition {
  display: flex;
  flex-direction: column;
  justify-content: flex-end;
  height: 100%;
  background-color: #000;
  color: #fff;
  font-size: 2.000000000000001rem;
  margin-left: 0.25rem;
  padding-left: 0.25rem;
  padding-right: 0.25rem;
}
.lcars-element .lcars-element-decorator {
  display: inline-block;
  position: absolute;
  top: 0;
  width: 2rem;
  left: -2.25rem;
  height: 100%;
  background-color: #fc6;
  border-top-left-radius: 1.5rem;
  border-bottom-left-radius: 1.5rem;
}
.lcars-element .lcars-element-decorator:after {
  content: '';
  display: block;
  position: absolute;
  right: 0.25rem;
  width: 0.25rem;
  height: 100%;
  background-color: #000;
}
.lcars-element .lcars-element-decorator.right {
  right: -2.25rem;
  left: unset;
  border-top-right-radius: 1.5rem;
  border-bottom-right-radius: 1.5rem;
  border-top-left-radius: 0;
  border-bottom-left-radius: 0;
}
.lcars-element .lcars-element-decorator.right:after {
  left: 0.25rem;
  right: unset;
}
.lcars-text-box,
a.lcars-text-box {
  background-color: #000;
  color: #fff;
  font-weight: bold;
  width: 7.5rem;
  height: 3rem;
  padding-left: 0.25rem;
  padding-right: 0.25rem;
  box-sizing: border-box;
  display: inline-flex;
  flex-direction: row;
}
.lcars-text-box.big,
a.lcars-text-box.big {
  font-size: 150%;
}
.lcars-text-box.large,
a.lcars-text-box.large {
  font-size: 200%;
}
.lcars-text-box.huge,
a.lcars-text-box.huge {
  font-size: 400%;
  height: 6rem;
}
.lcars-text-box.centered,
a.lcars-text-box.centered {
  justify-content: center;
}
.lcars-text-box.right,
a.lcars-text-box.right {
  justify-content: flex-end;
}
.lcars-text-box.full-centered,
a.lcars-text-box.full-centered {
  justify-content: center;
  align-items: center;
}
.lcars-text-box.centered-right,
a.lcars-text-box.centered-right {
  justify-content: flex-end;
  align-items: center;
}
.lcars-text-box.bottom,
a.lcars-text-box.bottom {
  align-items: flex-end;
}
.lcars-text-box.bottom-centered,
a.lcars-text-box.bottom-centered {
  justify-content: center;
  align-items: flex-end;
}
.lcars-text-box.bottom-right,
a.lcars-text-box.bottom-right {
  justify-content: flex-end;
  align-items: flex-end;
}
.lcars-text-box.blink,
a.lcars-text-box.blink {
  animation: blinker 1s linear infinite;
}
@-moz-keyframes blinker {
  50% {
    opacity: 0.4;
  }
}
@-webkit-keyframes blinker {
  50% {
    opacity: 0.4;
  }
}
@-o-keyframes blinker {
  50% {
    opacity: 0.4;
  }
}
@keyframes blinker {
  50% {
    opacity: 0.4;
  }
}
.lcars-bracket {
  background-color: #fc6;
}
.lcars-bracket.left {
  width: 1.5rem;
  border-top-left-radius: 1.5rem;
  border-bottom-left-radius: 1.5rem;
  border-width: 0.75rem 0 0.75rem 0.25rem;
}
.lcars-bracket.right {
  width: 1.5rem;
  border-top-right-radius: 1.5rem;
  border-bottom-right-radius: 1.5rem;
  border-width: 0.75rem 0.25rem 0.75rem 0;
}
.lcars-bracket.top {
  height: 1.5rem;
  border-top-left-radius: 1.5rem;
  border-top-right-radius: 1.5rem;
  border-width: 0.25rem 0.75rem 0 0.75rem;
}
.lcars-bracket.bottom {
  height: 1.5rem;
  border-bottom-left-radius: 1.5rem;
  border-bottom-right-radius: 1.5rem;
  border-width: 0 0.75rem 0.25rem 0.75rem;
}
.lcars-bracket.hollow {
  background-color: #000;
  border-style: solid;
  border-color: #fc6;
}
.lcars-u-1-1 {
  width: 7.5rem;
  min-width: 7.5rem;
  height: 3rem;
  min-height: 3rem;
}
.lcars-u-1-2 {
  width: 7.5rem;
  min-width: 7.5rem;
  height: 6.25rem;
  min-height: 6.25rem;
}
.lcars-u-1-3 {
  width: 7.5rem;
  min-width: 7.5rem;
  height: 9.5rem;
  min-height: 9.5rem;
}
.lcars-u-1-4 {
  width: 7.5rem;
  min-width: 7.5rem;
  height: 12.75rem;
  min-height: 12.75rem;
}
.lcars-u-1-5 {
  width: 7.5rem;
  min-width: 7.5rem;
  height: 16rem;
  min-height: 16rem;
}
.lcars-u-1-6 {
  width: 7.5rem;
  min-width: 7.5rem;
  height: 19.25rem;
  min-height: 19.25rem;
}
.lcars-u-1-7 {
  width: 7.5rem;
  min-width: 7.5rem;
  height: 22.5rem;
  min-height: 22.5rem;
}
.lcars-u-1-8 {
  width: 7.5rem;
  min-width: 7.5rem;
  height: 25.75rem;
  min-height: 25.75rem;
}
.lcars-u-2-1 {
  width: 15.25rem;
  min-width: 15.25rem;
  height: 3rem;
  min-height: 3rem;
}
.lcars-u-2-2 {
  width: 15.25rem;
  min-width: 15.25rem;
  height: 6.25rem;
  min-height: 6.25rem;
}
.lcars-u-2-3 {
  width: 15.25rem;
  min-width: 15.25rem;
  height: 9.5rem;
  min-height: 9.5rem;
}
.lcars-u-2-4 {
  width: 15.25rem;
  min-width: 15.25rem;
  height: 12.75rem;
  min-height: 12.75rem;
}
.lcars-u-2-5 {
  width: 15.25rem;
  min-width: 15.25rem;
  height: 16rem;
  min-height: 16rem;
}
.lcars-u-2-6 {
  width: 15.25rem;
  min-width: 15.25rem;
  height: 19.25rem;
  min-height: 19.25rem;
}
.lcars-u-2-7 {
  width: 15.25rem;
  min-width: 15.25rem;
  height: 22.5rem;
  min-height: 22.5rem;
}
.lcars-u-2-8 {
  width: 15.25rem;
  min-width: 15.25rem;
  height: 25.75rem;
  min-height: 25.75rem;
}
.lcars-u-3-1 {
  width: 23rem;
  min-width: 23rem;
  height: 3rem;
  min-height: 3rem;
}
.lcars-u-3-2 {
  width: 23rem;
  min-width: 23rem;
  height: 6.25rem;
  min-height: 6.25rem;
}
.lcars-u-3-3 {
  width: 23rem;
  min-width: 23rem;
  height: 9.5rem;
  min-height: 9.5rem;
}
.lcars-u-3-4 {
  width: 23rem;
  min-width: 23rem;
  height: 12.75rem;
  min-height: 12.75rem;
}
.lcars-u-3-5 {
  width: 23rem;
  min-width: 23rem;
  height: 16rem;
  min-height: 16rem;
}
.lcars-u-3-6 {
  width: 23rem;
  min-width: 23rem;
  height: 19.25rem;
  min-height: 19.25rem;
}
.lcars-u-3-7 {
  width: 23rem;
  min-width: 23rem;
  height: 22.5rem;
  min-height: 22.5rem;
}
.lcars-u-3-8 {
  width: 23rem;
  min-width: 23rem;
  height: 25.75rem;
  min-height: 25.75rem;
}
.lcars-u-4-1 {
  width: 30.75rem;
  min-width: 30.75rem;
  height: 3rem;
  min-height: 3rem;
}
.lcars-u-4-2 {
  width: 30.75rem;
  min-width: 30.75rem;
  height: 6.25rem;
  min-height: 6.25rem;
}
.lcars-u-4-3 {
  width: 30.75rem;
  min-width: 30.75rem;
  height: 9.5rem;
  min-height: 9.5rem;
}
.lcars-u-4-4 {
  width: 30.75rem;
  min-width: 30.75rem;
  height: 12.75rem;
  min-height: 12.75rem;
}
.lcars-u-4-5 {
  width: 30.75rem;
  min-width: 30.75rem;
  height: 16rem;
  min-height: 16rem;
}
.lcars-u-4-6 {
  width: 30.75rem;
  min-width: 30.75rem;
  height: 19.25rem;
  min-height: 19.25rem;
}
.lcars-u-4-7 {
  width: 30.75rem;
  min-width: 30.75rem;
  height: 22.5rem;
  min-height: 22.5rem;
}
.lcars-u-4-8 {
  width: 30.75rem;
  min-width: 30.75rem;
  height: 25.75rem;
  min-height: 25.75rem;
}
.lcars-u-5-1 {
  width: 38.5rem;
  min-width: 38.5rem;
  height: 3rem;
  min-height: 3rem;
}
.lcars-u-5-2 {
  width: 38.5rem;
  min-width: 38.5rem;
  height: 6.25rem;
  min-height: 6.25rem;
}
.lcars-u-5-3 {
  width: 38.5rem;
  min-width: 38.5rem;
  height: 9.5rem;
  min-height: 9.5rem;
}
.lcars-u-5-4 {
  width: 38.5rem;
  min-width: 38.5rem;
  height: 12.75rem;
  min-height: 12.75rem;
}
.lcars-u-5-5 {
  width: 38.5rem;
  min-width: 38.5rem;
  height: 16rem;
  min-height: 16rem;
}
.lcars-u-5-6 {
  width: 38.5rem;
  min-width: 38.5rem;
  height: 19.25rem;
  min-height: 19.25rem;
}
.lcars-u-5-7 {
  width: 38.5rem;
  min-width: 38.5rem;
  height: 22.5rem;
  min-height: 22.5rem;
}
.lcars-u-5-8 {
  width: 38.5rem;
  min-width: 38.5rem;
  height: 25.75rem;
  min-height: 25.75rem;
}
.lcars-u-6-1 {
  width: 46.25rem;
  min-width: 46.25rem;
  height: 3rem;
  min-height: 3rem;
}
.lcars-u-6-2 {
  width: 46.25rem;
  min-width: 46.25rem;
  height: 6.25rem;
  min-height: 6.25rem;
}
.lcars-u-6-3 {
  width: 46.25rem;
  min-width: 46.25rem;
  height: 9.5rem;
  min-height: 9.5rem;
}
.lcars-u-6-4 {
  width: 46.25rem;
  min-width: 46.25rem;
  height: 12.75rem;
  min-height: 12.75rem;
}
.lcars-u-6-5 {
  width: 46.25rem;
  min-width: 46.25rem;
  height: 16rem;
  min-height: 16rem;
}
.lcars-u-6-6 {
  width: 46.25rem;
  min-width: 46.25rem;
  height: 19.25rem;
  min-height: 19.25rem;
}
.lcars-u-6-7 {
  width: 46.25rem;
  min-width: 46.25rem;
  height: 22.5rem;
  min-height: 22.5rem;
}
.lcars-u-6-8 {
  width: 46.25rem;
  min-width: 46.25rem;
  height: 25.75rem;
  min-height: 25.75rem;
}
.lcars-u-7-1 {
  width: 54rem;
  min-width: 54rem;
  height: 3rem;
  min-height: 3rem;
}
.lcars-u-7-2 {
  width: 54rem;
  min-width: 54rem;
  height: 6.25rem;
  min-height: 6.25rem;
}
.lcars-u-7-3 {
  width: 54rem;
  min-width: 54rem;
  height: 9.5rem;
  min-height: 9.5rem;
}
.lcars-u-7-4 {
  width: 54rem;
  min-width: 54rem;
  height: 12.75rem;
  min-height: 12.75rem;
}
.lcars-u-7-5 {
  width: 54rem;
  min-width: 54rem;
  height: 16rem;
  min-height: 16rem;
}
.lcars-u-7-6 {
  width: 54rem;
  min-width: 54rem;
  height: 19.25rem;
  min-height: 19.25rem;
}
.lcars-u-7-7 {
  width: 54rem;
  min-width: 54rem;
  height: 22.5rem;
  min-height: 22.5rem;
}
.lcars-u-7-8 {
  width: 54rem;
  min-width: 54rem;
  height: 25.75rem;
  min-height: 25.75rem;
}
.lcars-u-8-1 {
  width: 61.75rem;
  min-width: 61.75rem;
  height: 3rem;
  min-height: 3rem;
}
.lcars-u-8-2 {
  width: 61.75rem;
  min-width: 61.75rem;
  height: 6.25rem;
  min-height: 6.25rem;
}
.lcars-u-8-3 {
  width: 61.75rem;
  min-width: 61.75rem;
  height: 9.5rem;
  min-height: 9.5rem;
}
.lcars-u-8-4 {
  width: 61.75rem;
  min-width: 61.75rem;
  height: 12.75rem;
  min-height: 12.75rem;
}
.lcars-u-8-5 {
  width: 61.75rem;
  min-width: 61.75rem;
  height: 16rem;
  min-height: 16rem;
}
.lcars-u-8-6 {
  width: 61.75rem;
  min-width: 61.75rem;
  height: 19.25rem;
  min-height: 19.25rem;
}
.lcars-u-8-7 {
  width: 61.75rem;
  min-width: 61.75rem;
  height: 22.5rem;
  min-height: 22.5rem;
}
.lcars-u-8-8 {
  width: 61.75rem;
  min-width: 61.75rem;
  height: 25.75rem;
  min-height: 25.75rem;
}
.lcars-u-9-1 {
  width: 69.5rem;
  min-width: 69.5rem;
  height: 3rem;
  min-height: 3rem;
}
.lcars-u-9-2 {
  width: 69.5rem;
  min-width: 69.5rem;
  height: 6.25rem;
  min-height: 6.25rem;
}
.lcars-u-9-3 {
  width: 69.5rem;
  min-width: 69.5rem;
  height: 9.5rem;
  min-height: 9.5rem;
}
.lcars-u-9-4 {
  width: 69.5rem;
  min-width: 69.5rem;
  height: 12.75rem;
  min-height: 12.75rem;
}
.lcars-u-9-5 {
  width: 69.5rem;
  min-width: 69.5rem;
  height: 16rem;
  min-height: 16rem;
}
.lcars-u-9-6 {
  width: 69.5rem;
  min-width: 69.5rem;
  height: 19.25rem;
  min-height: 19.25rem;
}
.lcars-u-9-7 {
  width: 69.5rem;
  min-width: 69.5rem;
  height: 22.5rem;
  min-height: 22.5rem;
}
.lcars-u-9-8 {
  width: 69.5rem;
  min-width: 69.5rem;
  height: 25.75rem;
  min-height: 25.75rem;
}
.lcars-u-10-1 {
  width: 77.25rem;
  min-width: 77.25rem;
  height: 3rem;
  min-height: 3rem;
}
.lcars-u-10-2 {
  width: 77.25rem;
  min-width: 77.25rem;
  height: 6.25rem;
  min-height: 6.25rem;
}
.lcars-u-10-3 {
  width: 77.25rem;
  min-width: 77.25rem;
  height: 9.5rem;
  min-height: 9.5rem;
}
.lcars-u-10-4 {
  width: 77.25rem;
  min-width: 77.25rem;
  height: 12.75rem;
  min-height: 12.75rem;
}
.lcars-u-10-5 {
  width: 77.25rem;
  min-width: 77.25rem;
  height: 16rem;
  min-height: 16rem;
}
.lcars-u-10-6 {
  width: 77.25rem;
  min-width: 77.25rem;
  height: 19.25rem;
  min-height: 19.25rem;
}
.lcars-u-10-7 {
  width: 77.25rem;
  min-width: 77.25rem;
  height: 22.5rem;
  min-height: 22.5rem;
}
.lcars-u-10-8 {
  width: 77.25rem;
  min-width: 77.25rem;
  height: 25.75rem;
  min-height: 25.75rem;
}
.lcars-u-11-1 {
  width: 85rem;
  min-width: 85rem;
  height: 3rem;
  min-height: 3rem;
}
.lcars-u-11-2 {
  width: 85rem;
  min-width: 85rem;
  height: 6.25rem;
  min-height: 6.25rem;
}
.lcars-u-11-3 {
  width: 85rem;
  min-width: 85rem;
  height: 9.5rem;
  min-height: 9.5rem;
}
.lcars-u-11-4 {
  width: 85rem;
  min-width: 85rem;
  height: 12.75rem;
  min-height: 12.75rem;
}
.lcars-u-11-5 {
  width: 85rem;
  min-width: 85rem;
  height: 16rem;
  min-height: 16rem;
}
.lcars-u-11-6 {
  width: 85rem;
  min-width: 85rem;
  height: 19.25rem;
  min-height: 19.25rem;
}
.lcars-u-11-7 {
  width: 85rem;
  min-width: 85rem;
  height: 22.5rem;
  min-height: 22.5rem;
}
.lcars-u-11-8 {
  width: 85rem;
  min-width: 85rem;
  height: 25.75rem;
  min-height: 25.75rem;
}
.lcars-u-12-1 {
  width: 92.75rem;
  min-width: 92.75rem;
  height: 3rem;
  min-height: 3rem;
}
.lcars-u-12-2 {
  width: 92.75rem;
  min-width: 92.75rem;
  height: 6.25rem;
  min-height: 6.25rem;
}
.lcars-u-12-3 {
  width: 92.75rem;
  min-width: 92.75rem;
  height: 9.5rem;
  min-height: 9.5rem;
}
.lcars-u-12-4 {
  width: 92.75rem;
  min-width: 92.75rem;
  height: 12.75rem;
  min-height: 12.75rem;
}
.lcars-u-12-5 {
  width: 92.75rem;
  min-width: 92.75rem;
  height: 16rem;
  min-height: 16rem;
}
.lcars-u-12-6 {
  width: 92.75rem;
  min-width: 92.75rem;
  height: 19.25rem;
  min-height: 19.25rem;
}
.lcars-u-12-7 {
  width: 92.75rem;
  min-width: 92.75rem;
  height: 22.5rem;
  min-height: 22.5rem;
}
.lcars-u-12-8 {
  width: 92.75rem;
  min-width: 92.75rem;
  height: 25.75rem;
  min-height: 25.75rem;
}
.lcars-u-13-1 {
  width: 100.5rem;
  min-width: 100.5rem;
  height: 3rem;
  min-height: 3rem;
}
.lcars-u-13-2 {
  width: 100.5rem;
  min-width: 100.5rem;
  height: 6.25rem;
  min-height: 6.25rem;
}
.lcars-u-13-3 {
  width: 100.5rem;
  min-width: 100.5rem;
  height: 9.5rem;
  min-height: 9.5rem;
}
.lcars-u-13-4 {
  width: 100.5rem;
  min-width: 100.5rem;
  height: 12.75rem;
  min-height: 12.75rem;
}
.lcars-u-13-5 {
  width: 100.5rem;
  min-width: 100.5rem;
  height: 16rem;
  min-height: 16rem;
}
.lcars-u-13-6 {
  width: 100.5rem;
  min-width: 100.5rem;
  height: 19.25rem;
  min-height: 19.25rem;
}
.lcars-u-13-7 {
  width: 100.5rem;
  min-width: 100.5rem;
  height: 22.5rem;
  min-height: 22.5rem;
}
.lcars-u-13-8 {
  width: 100.5rem;
  min-width: 100.5rem;
  height: 25.75rem;
  min-height: 25.75rem;
}
.lcars-u-14-1 {
  width: 108.25rem;
  min-width: 108.25rem;
  height: 3rem;
  min-height: 3rem;
}
.lcars-u-14-2 {
  width: 108.25rem;
  min-width: 108.25rem;
  height: 6.25rem;
  min-height: 6.25rem;
}
.lcars-u-14-3 {
  width: 108.25rem;
  min-width: 108.25rem;
  height: 9.5rem;
  min-height: 9.5rem;
}
.lcars-u-14-4 {
  width: 108.25rem;
  min-width: 108.25rem;
  height: 12.75rem;
  min-height: 12.75rem;
}
.lcars-u-14-5 {
  width: 108.25rem;
  min-width: 108.25rem;
  height: 16rem;
  min-height: 16rem;
}
.lcars-u-14-6 {
  width: 108.25rem;
  min-width: 108.25rem;
  height: 19.25rem;
  min-height: 19.25rem;
}
.lcars-u-14-7 {
  width: 108.25rem;
  min-width: 108.25rem;
  height: 22.5rem;
  min-height: 22.5rem;
}
.lcars-u-14-8 {
  width: 108.25rem;
  min-width: 108.25rem;
  height: 25.75rem;
  min-height: 25.75rem;
}
.lcars-u-15-1 {
  width: 116rem;
  min-width: 116rem;
  height: 3rem;
  min-height: 3rem;
}
.lcars-u-15-2 {
  width: 116rem;
  min-width: 116rem;
  height: 6.25rem;
  min-height: 6.25rem;
}
.lcars-u-15-3 {
  width: 116rem;
  min-width: 116rem;
  height: 9.5rem;
  min-height: 9.5rem;
}
.lcars-u-15-4 {
  width: 116rem;
  min-width: 116rem;
  height: 12.75rem;
  min-height: 12.75rem;
}
.lcars-u-15-5 {
  width: 116rem;
  min-width: 116rem;
  height: 16rem;
  min-height: 16rem;
}
.lcars-u-15-6 {
  width: 116rem;
  min-width: 116rem;
  height: 19.25rem;
  min-height: 19.25rem;
}
.lcars-u-15-7 {
  width: 116rem;
  min-width: 116rem;
  height: 22.5rem;
  min-height: 22.5rem;
}
.lcars-u-15-8 {
  width: 116rem;
  min-width: 116rem;
  height: 25.75rem;
  min-height: 25.75rem;
}
.lcars-u-16-1 {
  width: 123.75rem;
  min-width: 123.75rem;
  height: 3rem;
  min-height: 3rem;
}
.lcars-u-16-2 {
  width: 123.75rem;
  min-width: 123.75rem;
  height: 6.25rem;
  min-height: 6.25rem;
}
.lcars-u-16-3 {
  width: 123.75rem;
  min-width: 123.75rem;
  height: 9.5rem;
  min-height: 9.5rem;
}
.lcars-u-16-4 {
  width: 123.75rem;
  min-width: 123.75rem;
  height: 12.75rem;
  min-height: 12.75rem;
}
.lcars-u-16-5 {
  width: 123.75rem;
  min-width: 123.75rem;
  height: 16rem;
  min-height: 16rem;
}
.lcars-u-16-6 {
  width: 123.75rem;
  min-width: 123.75rem;
  height: 19.25rem;
  min-height: 19.25rem;
}
.lcars-u-16-7 {
  width: 123.75rem;
  min-width: 123.75rem;
  height: 22.5rem;
  min-height: 22.5rem;
}
.lcars-u-16-8 {
  width: 123.75rem;
  min-width: 123.75rem;
  height: 25.75rem;
  min-height: 25.75rem;
}
.lcars-vu-1 {
  height: 3rem;
}
.lcars-vu-2 {
  height: 6.25rem;
}
.lcars-vu-3 {
  height: 9.5rem;
}
.lcars-vu-4 {
  height: 12.75rem;
}
.lcars-vu-5 {
  height: 16rem;
}
.lcars-vu-6 {
  height: 19.25rem;
}
.lcars-vu-7 {
  height: 22.5rem;
}
.lcars-vu-8 {
  height: 25.75rem;
}
.lcars-vu-9 {
  height: 29rem;
}
.lcars-vu-10 {
  height: 32.25rem;
}
.lcars-vu-11 {
  height: 35.5rem;
}
.lcars-vu-12 {
  height: 38.75rem;
}
.lcars-vu-13 {
  height: 42rem;
}
.lcars-vu-14 {
  height: 45.25rem;
}
.lcars-vu-15 {
  height: 48.5rem;
}
.lcars-vu-16 {
  height: 51.75rem;
}
.lcars-u-1 {
  width: 7.5rem;
}
.lcars-u-1.half {
  width: 3.625rem;
}
.lcars-u-2 {
  width: 15.25rem;
}
.lcars-u-2.half {
  width: 11.375rem;
}
.lcars-u-3 {
  width: 23rem;
}
.lcars-u-3.half {
  width: 19.125rem;
}
.lcars-u-4 {
  width: 30.75rem;
}
.lcars-u-4.half {
  width: 26.875rem;
}
.lcars-u-5 {
  width: 38.5rem;
}
.lcars-u-5.half {
  width: 34.625rem;
}
.lcars-u-6 {
  width: 46.25rem;
}
.lcars-u-6.half {
  width: 42.375rem;
}
.lcars-u-7 {
  width: 54rem;
}
.lcars-u-7.half {
  width: 50.125rem;
}
.lcars-u-8 {
  width: 61.75rem;
}
.lcars-u-8.half {
  width: 57.875rem;
}
.lcars-u-9 {
  width: 69.5rem;
}
.lcars-u-9.half {
  width: 65.625rem;
}
.lcars-u-10 {
  width: 77.25rem;
}
.lcars-u-10.half {
  width: 73.375rem;
}
.lcars-u-11 {
  width: 85rem;
}
.lcars-u-11.half {
  width: 81.125rem;
}
.lcars-u-12 {
  width: 92.75rem;
}
.lcars-u-12.half {
  width: 88.875rem;
}
.lcars-u-13 {
  width: 100.5rem;
}
.lcars-u-13.half {
  width: 96.625rem;
}
.lcars-u-14 {
  width: 108.25rem;
}
.lcars-u-14.half {
  width: 104.375rem;
}
.lcars-u-15 {
  width: 116rem;
}
.lcars-u-15.half {
  width: 112.125rem;
}
.lcars-u-16 {
  width: 123.75rem;
}
.lcars-u-16.half {
  width: 119.875rem;
}
`
