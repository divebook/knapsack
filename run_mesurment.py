import subprocess
import os
import configparser
import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt


ABS_PATH = os.path.abspath(os.path.dirname(__file__))


def run_instance(num_repeat, optimized, algorithm_name, num_items, num_instances, capacity_weights_ratio, max_weight, max_price, exponent, type):

    df = pd.DataFrame()

    for i in range(num_repeat):

        subprocess.call(['./generator.sh',
                     '-o', f'{optimized}',
                     '-a', f'{algorithm_name}',
                     '-n', f'{num_items}',
                     '-N', f'{num_instances}',
                     '-m', f'{capacity_weights_ratio}',
                     '-W', f'{max_weight}',
                     '-C', f'{max_price}',
                     '-k', f'{exponent}',
                     '-d', f'{type}'])

        new_df = pd.read_csv('instance_data.csv')

        if i != 0:
            df['duration'] = df['duration'] + new_df['duration']
            if algorithm_name == 'heuristic':
                df['error'] = df['error'] + new_df['error']
        else:
            df = new_df



    df['duration'] = df['duration']/num_repeat

    if algorithm_name == 'heuristic':
        df['error'] = df['error']/num_repeat


    print(df.head())

    new_filename = 'output/' + 'inst'+ '_' + str(optimized) + '_' + str(algorithm_name) + '_' + str(capacity_weights_ratio) + '_' + str(exponent)

    sns.set(style="darkgrid")
    sns_plot = sns.lineplot(x=optimized, y='duration', data=df).get_figure()
    sns_plot.savefig(new_filename + '.png')

    if algorithm_name == 'heuristic':
        plt.clf()
        sns_plot = sns.lineplot(x=optimized, y='error', data=df).get_figure()
        sns_plot.savefig(new_filename + '_error' + '.png')

    # clear plot
    plt.clf()

    df.to_csv(new_filename + '.csv', sep=',')





if __name__ == "__main__":

    path = ABS_PATH + '/config'
    #configs = [f for f in os.listdir(path) if os.path.isfile(os.path.join(path, f))]

    #configs = ['config_bb_weight.cfg', 'config_bb_price.cfg', 'config_bb_ratio.cfg', 'config_bb_exp.cfg']
    configs = ['config_bb_exp.cfg', 'config_dynamic_price_exp.cfg', 'config_dynamic_weight_exp.cfg', 'config_heuristic_exp.cfg']
    #configs = ['config_dynamic_price_price.cfg']

    for i in range(len(configs)):

        print(ABS_PATH + '/config/' + configs[i])

        cfg = configparser.ConfigParser()
        cfg.read(ABS_PATH + '/config/' + configs[i])

        num_repeat = cfg.getint('params', 'num_repeat')
        optimized = cfg.get('params', 'optimized')
        algorithm_name = cfg.get('params', 'algorithm_name')
        num_items = cfg.getint('params', 'num_items')
        num_instances = cfg.getint('params', 'num_instances')
        capacity_weights_ratio = cfg.getfloat('params', 'capacity_weights_ratio')
        max_weight = cfg.getint('params', 'max_weight')
        max_price = cfg.getint('params', 'max_price')
        exponent = cfg.getfloat('params', 'exponent')
        type = cfg.getint('params', 'type')

        run_instance(num_repeat, optimized, algorithm_name, num_items, num_instances, capacity_weights_ratio, max_weight, max_price, exponent, type)