// @flow
import * as React from 'react'
import {globalStyles, globalColors} from '../../styles'
import {Text} from '../../common-adapters'
import {type UploadProps} from './upload'
import {NativeAnimated} from '../../common-adapters/native-wrappers.native'

const patternURL = require('../../images/upload-pattern-80.png')

type UploadState = {
  backgroundPositionY: NativeAnimated.AnimatedValue,
}

class Upload extends React.PureComponent<UploadProps, UploadState> {
  state = {
    backgroundPositionY: new NativeAnimated.Value(0),
  }

  componentDidMount() {
    NativeAnimated.timing(this.state.backgroundPositionY, {
      toValue: -48,
      duration: 2000,
      useNativeDriver: true,
    }).loop()
  }

  render() {
    const {backgroundPositionY} = this.state
    const {files, timeLeft} = this.props
    return (
      <NativeAnimated.View style={{...stylesBox, backgroundPositionY}}>
        <Text
          key="files"
          type="BodySemibold"
          style={stylesText}
        >{`Encrypting and uploading ${files} files...`}</Text>
        {!!(timeLeft && timeLeft.length) && (
          <Text key="left" type="BodySmall" style={stylesText}>{`${timeLeft} left`}</Text>
        )}
      </NativeAnimated.View>
    )
  }
}

const stylesText = {
  color: globalColors.white,
}

const stylesBox = {
  ...globalStyles.flexBoxColumn,
  alignItems: 'center',
  justifyContent: 'center',
  backgroundImage: `url('${patternURL}')`,
  height: 48,
}

export default Upload
