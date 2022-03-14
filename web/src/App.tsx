import { Theme, ThemeProvider } from '@fluentui/react';
import React, { useEffect, useMemo, useState } from 'react';
import { I18nextProvider } from 'react-i18next';
import './App.css';
import { ErrorMessageBar } from './components/common/ErrorMessageBar';
import { Loading } from './components/common/Loading';
import { MainTab } from './components/MainTab';
import { ConfigEvent, ConfigEventAction, IConfigEvent } from './events/ConfigEvent';
import i18n from './i18n';
import './RegisterIcons';
import { Config, getConfig, setConfig } from './services/config.service';
import { getCustomTheme } from './themes';

function App() {
  const [_config, _setConfig] = useState<Config>({ theme: 'light', language: navigator.language}),
    [theme, setTheme] = useState<Theme>(getCustomTheme('light')),
    [error, setError] = useState<Error>(),
    [loading, setLoading] = useState<boolean>(true);

  useEffect(() => {
    setLoading(true);
    getConfig().then((ret) => {
      if (!ret) return;
      _setConfig({ ...ret });
      setTheme(getCustomTheme(ret.theme));
    })
      .catch(setError)
      .finally(() => {
        setLoading(false);
      })
  }, []);

  useEffect(() => {
    const sub = ConfigEvent.subscribe((v: IConfigEvent) => {
      switch (v.action) {
        case ConfigEventAction.Language:
          _setConfig(c => {
            c.language = v.params;
            saveConfig(c);
            return c;
          })
          i18n.changeLanguage(v.params);
          break;
        case ConfigEventAction.Theme:
          _setConfig(c => {
            c.theme = v.params;
            saveConfig(c);
            return c;
          })
          setTheme(getCustomTheme(v.params));
          break;
        default:
          break;
      }
    })
    return () => sub && sub.unsubscribe();
  }, [_config])

  const saveConfig = (conf: Config) => {
    setLoading(true);
    setConfig(conf).catch(setError).finally(() => {
      setLoading(false);
    })
  }

  const mainTab = useMemo(() => (
    <MainTab config={_config}></MainTab>
  ), [_config])

  return (<>
    <ThemeProvider theme={theme} style={{ height: '100%' }}>
      <I18nextProvider i18n={i18n}>
        {/* loading */}
        <Loading loading={loading} />
        {/* error */}
        <ErrorMessageBar error={error}></ErrorMessageBar>
        {/* mainTab */}
        {mainTab}
        {/* <MainTab config={_config}
          onChangeLanguage={handleChnageLanguage}
          onChangeTheme={handleChangeTheme}
          onChangePort={handleChangePort}
        ></MainTab> */}
      </I18nextProvider>
    </ThemeProvider>
  </>);
}

export default App;
