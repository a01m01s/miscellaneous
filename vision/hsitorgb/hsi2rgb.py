# python implementation of the matlab code:
# https://stackoverflow.com/a/74061828

import sys

import imageio
import numpy as np
import scipy.io as spio


def main():
    try:
        HSI_mat = spio.loadmat(sys.argv[1])
        HSI_data = HSI_mat[sys.argv[2]]

        red = HSI_data[:, :, 53]
        green = HSI_data[:, :, 21]
        blue = HSI_data[:, :, 7]

        rgb_image = np.stack((red, green, blue), axis=2)
        rgb_image = rgb_image - rgb_image.min()
        rgb_image = ((rgb_image/rgb_image.max())*255.0).astype(np.uint8)

        imageio.imwrite('output.png', rgb_image)

    except:
        e = sys.exc_info()[0]
        print(e)


if __name__ == '__main__':
    main()
