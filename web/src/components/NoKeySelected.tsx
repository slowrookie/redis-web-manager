import { FontIcon, Stack, Text, Link, useTheme } from '@fluentui/react'
import React from 'react'
import { useTranslation } from 'react-i18next'


const NoKeySelected = (props: any) => {
  const theme = useTheme(), { t } = useTranslation();

  const fontColorStyle = {
    color: theme.palette.neutralTertiaryAlt
  }

  return (
    <Stack horizontalAlign="center" tokens={{ childrenGap: 10 }}
      style={{ height: '100%', paddingTop: '10%' }}>
      <FontIcon iconName="Emoji2" style={{ fontSize: 64, ...fontColorStyle }}></FontIcon>
      <Link href="https://github.com/slowrookie/redis-web-manager/discussions" target="_blank">
        <Text variant="small" style={fontColorStyle}>{t('Welcome to discussions!')}</Text>
      </Link>
      <Link href="https://github.com/slowrookie/redis-web-manager" target="_blank" style={fontColorStyle}>
        https://github.com/slowrookie/redis-web-manager
      </Link>
    </Stack>
  )
}

export default NoKeySelected
