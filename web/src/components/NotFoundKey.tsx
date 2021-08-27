import {FontIcon, Stack, Text, useTheme} from '@fluentui/react'
import React from 'react'

export interface INotFoundProps {
  message?: string
}

const NotFoundKey = (props: INotFoundProps) => {
  const theme = useTheme()

  return (
    <Stack horizontalAlign="center" tokens={{childrenGap: 10}} style={{height: '100%', paddingTop: '10%'}}>
      <FontIcon iconName="Sad" style={{fontSize: 64}}></FontIcon>
      <Text variant="xLarge"> 404 </Text>
      {props.message && (
        <Text variant="medium" style={{color: theme.palette.redDark}}>
          {props.message}
        </Text>
      )}
    </Stack>
  )
}

export default NotFoundKey
