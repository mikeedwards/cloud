
import React, { PropTypes } from 'react'
import { dateToString } from '../../../utils.js'
import PlaybackSelector from './PlaybackSelector'
import FocusSelector from './FocusSelector'
import ZoomSelector from './ZoomSelector'

class ControlPanel extends React.Component {

  render () {

    const { 
      currentExpeditionID,
      currentDate,
      playbackMode,
      focus,
      zoom,
      selectPlaybackMode,
      selectFocusType,
      selectZoom,
      toggleSensorData,
      showSensors
    } = this.props

    const currentPage = location.pathname.split('/').filter(p => !!p && p !== currentExpeditionID)[0] || 'map'

    const sensorToggleTitle = !showSensors ? "Social + data" : "Only social"
    const sensorToggleClass = showSensors ? "control-panel_toggle-sensor" : "control-panel_toggle-sensor off"
    return (
      <div className="control-panel">
        <div className="control-panel_date-counter">
          { dateToString(new Date(currentDate)) }
        </div>
        {
          currentPage === 'map' &&
          <div>
            <PlaybackSelector
              playbackMode={ playbackMode }
              selectPlaybackMode={ selectPlaybackMode }
            />
            <FocusSelector
              focus={ focus }
              selectFocusType={ selectFocusType }
            />
            <ZoomSelector
              zoom={ zoom }
              selectZoom={ selectZoom }
            />
            <div
                className={sensorToggleClass}
                onClick={ () => toggleSensorData() }
            > 
                {sensorToggleTitle} 
            </div>
          </div>
        }
      </div>
    )
  }

}

ControlPanel.propTypes = {

}

export default ControlPanel