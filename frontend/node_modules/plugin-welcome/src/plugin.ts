import { createPlugin } from '@backstage/core';
import WelcomePage from './components/WelcomePage';
import CreateDeposit from './components/RecordDeposit';
import ShowDeposit from './components/DepositTable';
import SignIn from './components/SignIn';


export const plugin = createPlugin({
  id: 'welcome',
  register({ router }) {
    router.registerRoute('/', SignIn);
	router.registerRoute('/Home', WelcomePage);
    router.registerRoute('/DepositTable', ShowDeposit);
    router.registerRoute('/RecordDeposit', CreateDeposit);
  },
});
